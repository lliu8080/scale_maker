package api

import (
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
	_ "nuc.lliu.ca/gitea/app/scale_maker/docs" // swagger doc
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
	v1.Get("/namespace/list", timeout.NewWithContext(listNamespaces, timeOut*time.Second))

	//node related APIs
	v1.Get("/node/list", timeout.NewWithContext(listNodes, timeOut*time.Second))

	// deployment related APIs
	v1.Get("/daemonset/list", timeout.NewWithContext(listDaemonsets, timeOut*time.Second))

	// deployment related APIs
	v1.Get("/deployment/list", timeout.NewWithContext(listDeployments, timeOut*time.Second))

	// pod related APIs
	v1.Get("/pod/list", timeout.NewWithContext(listPods, timeOut*time.Second))
	v1.Post("/pod/template/create", timeout.NewWithContext(createPodFromTemplate, timeOut*time.Second))
	v1.Post("/pod/yaml/create", timeout.NewWithContext(createPodFromBody, timeOut*time.Second))

	// job related APIs
	v1.Get("/job/list", timeout.NewWithContext(listJobs, timeOut*time.Second))
	v1.Post("/job/template/create", timeout.NewWithContext(createJobFromTemplate, timeOut*time.Second))
	v1.Post("/job/yaml/create", timeout.NewWithContext(createJobFromBody, timeOut*time.Second))

	// service related APIs
	v1.Get("/service/list", timeout.NewWithContext(listServices, timeOut*time.Second))

	// statefulset related APIs
	v1.Get("/statefulset/list", timeout.NewWithContext(listStatefulsets, timeOut*time.Second))

	// statefulset related APIs
	v1.Post("/bulk/create", timeout.NewWithContext(createResourcesFromBody, timeOut*time.Second))

	// Bind handlers
	v1.Get("/ping", timeout.NewWithContext(getStatus, timeOut*time.Second))

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
