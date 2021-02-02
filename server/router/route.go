package router

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/pascaloseko/shopping_list/server/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// routes
	api.Get("/", handler.GetAllItems)
	api.Get("/:id", handler.GetSingleItem)
	api.Post("/", handler.CreateItem)
	api.Delete("/:id", handler.DeleteItem)
	api.Put("/:id", handler.UpdateItem)
}