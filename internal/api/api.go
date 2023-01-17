package api

import (
	"context"
	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "nuc.lliu.ca/gitea/app/scale_maker/docs"
	"nuc.lliu.ca/gitea/app/scale_maker/internal/config"
	ctrl "sigs.k8s.io/controller-runtime"
)

type k8sClinet struct {
	clientSet     kubernetes.Interface //*kubernetes.Clientset or fake
	dynamicClient dynamic.Interface    //*dynamic.DynamicClient or fake
	ctx           context.Context
}

var k8sClients k8sClinet

func newK8SClient() {
	var err error
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	return nil, err
	// }

	// return kubernetes.NewForConfig(config)
	k8sClients.ctx = context.Background()
	config := ctrl.GetConfigOrDie()
	k8sClients.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error: unable to create normal Kubernetes clientSet.")
	}
	k8sClients.dynamicClient = dynamic.NewForConfigOrDie(config)
}

func setupRoutesandApp(app *fiber.App, testing bool) {

	//var err error
	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	// file, err := os.OpenFile("./item.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// pod related APIs
	v1.Get("/pod", ListPod)

	// namespace related APIs
	v1.Get("/namespace", ListNamespace)

	// Bind handlers
	v1.Get("/ping", GetStatus)

	app.Static("/favicon.ico", "./assets/static/img/favicon.ico")
	app.Get("/docs/*", swagger.HandlerDefault)

	// if err != nil {
	// 	log.Fatal("Error: cannot connect to the k8s cluster and initialize the client!")
	// }
	// app.Use(func(c *fiber.Ctx) error {
	// 	SetLocal[kubernetes.Interface](c, "k8s_client", k8s_client)
	// 	return c.Next()
	// })
	// Skip registering prometheus metrics if testing
	if !testing {
		prometheus := fiberprometheus.New("scale_maker")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
		app.Use(NotFound, recover.New())
		app.Use(logger.New())
	}
}

// InitialSetup doc
func InitialSetup() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})
	newK8SClient()
	setupRoutesandApp(app, false)

	// Return the configured app
	return app
}
