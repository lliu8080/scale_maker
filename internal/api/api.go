package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "nuc.lliu.ca/gitea/app/scale_maker/docs"
	"nuc.lliu.ca/gitea/app/scale_maker/internal/config"
)

func setupMiddleware(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
}

func setupRoutes(app *fiber.App) {
	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/ping", GetStatus)

	app.Static("/favicon.ico", "./assets/static/img/favicon.ico")
	app.Get("/docs/*", swagger.HandlerDefault)
}

func setupApp() *fiber.App {
	config.NewConfig()
	conf := config.AppConfig
	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: conf.Prod, // go run app.go -prod
	})

	setupMiddleware(app)
	setupRoutes(app)

	// Handle not founds
	app.Use(NotFound)

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
