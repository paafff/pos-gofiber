package handlers

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderedProductHandler struct {
	orderedProductService *services.OrderedProductService
}

func NewOrderedProductHandler(orderedProductService *services.OrderedProductService) *OrderedProductHandler {
	return &OrderedProductHandler{orderedProductService: orderedProductService}
}

func (h *OrderedProductHandler) CreateOrderedProduct(c *fiber.Ctx) error {

	orderedProduct := new(models.OrderedProduct)
	if err := c.BodyParser(orderedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	createdOrderedProduct, err := h.orderedProductService.CreateOrderedProduct(orderedProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create ordered product")
	}

	return c.JSON(createdOrderedProduct)
}

func (h *OrderedProductHandler) GetOrderedProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ordered product ID")
	}

	orderedProduct, err := h.orderedProductService.GetOrderedProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Ordered product not found")
	}

	return c.JSON(orderedProduct)
}

func (h *OrderedProductHandler) GetOrderedProducts(c *fiber.Ctx) error {
	orderedProducts, err := h.orderedProductService.GetOrderedProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve ordered products")
	}

	return c.JSON(orderedProducts)
}

func (h *OrderedProductHandler) UpdateOrderedProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ordered product ID")
	}

	orderedProduct := new(models.OrderedProduct)
	if err := c.BodyParser(orderedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}
	orderedProduct.ID = uint(id)

	updatedOrderedProduct, err := h.orderedProductService.UpdateOrderedProduct(orderedProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not update ordered product")
	}

	return c.JSON(updatedOrderedProduct)
}

func (h *OrderedProductHandler) DeleteOrderedProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ordered product ID")
	}

	err = h.orderedProductService.DeleteOrderedProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete ordered product")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
