package router

import (
	"bedtime-bully-backend/ent"
	"bedtime-bully-backend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, dbClient *ent.Client) {
	setUpWellKnownRoutes(app)
	setUpShareRedirectRoutes(app)

	api := app.Group("/api", logger.New())
	api.Get("/health", handler.HealthEndpoint())
	api.Group("/v1")
}
