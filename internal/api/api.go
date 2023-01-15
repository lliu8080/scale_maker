package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	_ "nuc.lliu.ca/gitea/app/scale_maker/docs"
	"nuc.lliu.ca/gitea/app/scale_maker/internal/config"
)

func newK8SClient() (kubernetes.Interface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)

}

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}

func setupRoutes(app *fiber.App) {

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	// file, err := os.OpenFile("./item.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	prometheus := fiberprometheus.New("scale_maker")
	prometheus.RegisterAt(app, "/metrics")

	// pod related APIs
	v1.Get("/pod", ListPod)

	// namespace related APIs
	v1.Get("/namespace", ListNamespace)

	// Bind handlers
	v1.Get("/ping", GetStatus)

	app.Static("/favicon.ico", "./assets/static/img/favicon.ico")
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Use(logger.New())
	app.Use(NotFound, recover.New(), prometheus.Middleware)
	k8s_client, _ := newK8SClient()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("k8s_client", k8s_client)
		return c.Next()
	})
}

func setupApp() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})

	setupRoutes(app)

	// Return the configured app
	return app
}

func InitialSetup() *fiber.App {
	// dbUser := "test"
	// dbPass := "test"
	// dbName := "weather_app"
	// dbHostName := "postgresql-weather_app_fiber"
	// dbPort := "5432"
	// config.InitDatabase(dbUser, dbPass, dbName, dbHostName, dbPort)
	return setupApp()
}
