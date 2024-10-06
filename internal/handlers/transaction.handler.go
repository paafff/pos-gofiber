package handlers

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	createdTransaction, err := h.transactionService.CreateTransaction(transaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create transaction")
	}

	return c.JSON(createdTransaction)
}

func (h *TransactionHandler) GetTransaction(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
	}

	transaction, err := h.transactionService.GetTransaction(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Transaction not found")
	}

	return c.JSON(transaction)
}

func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	transactions, err := h.transactionService.GetTransactions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve transactions")
	}

	return c.JSON(transactions)
}

func (h *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
	}

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}
	transaction.ID = uint(id)

	updatedTransaction, err := h.transactionService.UpdateTransaction(transaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not update transaction")
	}

	return c.JSON(updatedTransaction)
}

// func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {

// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
// 	}

// 	err = h.transactionService.DeleteTransaction(uint(id))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete transaction")
// 	}

// 	return c.SendStatus(fiber.StatusNoContent)
// }

func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid transaction ID")
	}

	if err := h.transactionService.DeleteTransaction(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete transaction")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

//=======================================================
