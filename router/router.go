package router

import (
	"github.com/firmanjs/go-restfull-with-fiber/handler"

	"github.com/firmanjs/go-restfull-with-fiber/middleware"

	"github.com/gofiber/fiber"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", middleware.AuthReq())

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)

}
