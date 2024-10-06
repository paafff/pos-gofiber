package routes

import (
	"pos-gofiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func OrderedProductRoutes(router fiber.Router, orderedProductHandler *handlers.OrderedProductHandler) {
	router.Get("/ordered-product/:id", orderedProductHandler.GetOrderedProduct)
	router.Post("/ordered-product", orderedProductHandler.CreateOrderedProduct)
	router.Put("/ordered-product/:id", orderedProductHandler.UpdateOrderedProduct)
	router.Delete("/ordered-product/:id", orderedProductHandler.DeleteOrderedProduct)
	router.Get("/ordered-products", orderedProductHandler.GetOrderedProducts)
}
