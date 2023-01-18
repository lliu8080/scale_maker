package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// namespace related APIs
	v1.Get("/namespaces/list", listNamespaces)

	//node related APIs
	v1.Get("/nodes/list", listNodes)

	// deployment related APIs
	v1.Get("/daemonsets/list", listDaemonsets)

	// deployment related APIs
	v1.Get("/deployments/list", listDeployments)

	// pod related APIs
	v1.Get("/pods/list", listPods)
	v1.Get("/pods/create", createPod)

	// service related APIs
	v1.Get("/services/list", listServices)

	// statefulset related APIs
	v1.Get("/statefulsets/list", listStatefulsets)

	// Bind handlers
	v1.Get("/ping", getStatus)

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
