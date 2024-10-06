package routes

import (
	"pos-gofiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(router fiber.Router, transactionHandler *handlers.TransactionHandler) {
	router.Get("/transaction/:id", transactionHandler.GetTransaction)
	router.Post("/transaction", transactionHandler.CreateTransaction)
	router.Put("/transaction/:id", transactionHandler.UpdateTransaction)
	router.Delete("/transaction/:id", transactionHandler.DeleteTransaction)
	router.Get("/transactions", transactionHandler.GetTransactions)
}
