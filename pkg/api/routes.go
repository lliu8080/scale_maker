package api

import (
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
)

func setupRoutesandMiddleware(app *fiber.App, testing bool) {
	// file, err := os.OpenFile("./item.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	//var err error

	app.Use(recover.New())
	app.Use(logger.New())

	timeOut := time.Duration(300) // requests timeout in 5 minutes

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// namespace related APIs
	v1.Get("/namespace/list", timeout.New(listNamespaces, timeOut*time.Second))

	//node related APIs
	v1.Get("/node/list", timeout.New(listNodes, timeOut*time.Second))

	// deployment related APIs
	v1.Get("/daemonset/list", timeout.New(listDaemonsets, timeOut*time.Second))

	// deployment related APIs
	v1.Get("/deployment/list", timeout.New(listDeployments, timeOut*time.Second))

	// pod related APIs
	v1.Get("/pod/list", timeout.New(listPods, timeOut*time.Second))
	v1.Post("/pod/template/create", timeout.New(createPodFromTemplate, timeOut*time.Second))
	v1.Post("/pod/yaml/create", timeout.New(createPodFromBody, timeOut*time.Second))

	// service related APIs
	v1.Get("/service/list", timeout.New(listServices, timeOut*time.Second))

	// statefulset related APIs
	v1.Get("/statefulset/list", timeout.New(listStatefulsets, timeOut*time.Second))

	// statefulset related APIs
	v1.Get("/bulk/create", timeout.New(createResourcesFromBody, timeOut*time.Second))

	// Bind handlers
	v1.Get("/ping", timeout.New(getStatus, timeOut*time.Second))

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
	}
	app.Use(NotFound)
}
