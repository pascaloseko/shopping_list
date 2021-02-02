package router

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/pascaloseko/shopping_list/server/handler"
	"github.com/pascaloseko/shopping_list/server/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Public
	public := app.Group("/public", logger.New())
	// Middleware
	api := app.Group("/api", logger.New(), middleware.Authz())

	public.Post("/signup", handler.Register)
	public.Post("/login", handler.LoginUser)

	// routes
	api.Get("/", handler.GetAllItems)
	api.Get("/:id", handler.GetSingleItem)
	api.Post("/", handler.CreateItem)
	api.Delete("/:id", handler.DeleteItem)
	api.Put("/:id", handler.UpdateItem)
}
