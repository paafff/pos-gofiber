package router

import (
	"pos-gofiber/internal/auth"
	"pos-gofiber/internal/handlers"
	"pos-gofiber/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler,
	authHandler *auth.AuthHandler,
	productHandler *handlers.ProductHandler,
	transactionHandler *handlers.TransactionHandler,

	orderedProductHandler *handlers.OrderedProductHandler) {
	api := app.Group("/api")

	// Public routes
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)

	// Protected routes
	api.Use(auth.AuthMiddleware)
	//user routes
	routes.UserRoutes(api, userHandler)
	//product routes
	routes.ProductRoutes(api, productHandler)
	//transaction routes
	routes.TransactionRoutes(api, transactionHandler)
	//ordered product routes
	routes.OrderedProductRoutes(api, orderedProductHandler)
}
