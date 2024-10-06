package handlers

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	product, err := h.productService.GetProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	return c.JSON(product)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve products")
	}

	return c.JSON(products)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	createdProduct, err := h.productService.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create product")
	}

	return c.JSON(createdProduct)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	product.ID = uint(id)
	updatedProduct, err := h.productService.UpdateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not update product")
	}

	return c.JSON(updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete product")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
