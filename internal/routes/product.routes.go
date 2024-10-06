package routes

import (
	"pos-gofiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router, productHandler *handlers.ProductHandler) {
	router.Get("/product/:id", productHandler.GetProduct)
	router.Post("/product", productHandler.CreateProduct)
	router.Put("/product/:id", productHandler.UpdateProduct)
	router.Delete("/product/:id", productHandler.DeleteProduct)
	router.Get("/products", productHandler.GetProducts)
}
