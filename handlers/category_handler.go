package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Index(c *fiber.Ctx) error {
	id := c.Params("id")
	search := c.Query("search", "")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, count, err := h.service.GetAll(id, search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"data":   data,
	})
}

func (h *CategoryHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}

	return c.JSON(fiber.Map{"status": "success", "data": data})
}

func (h *CategoryHandler) Store(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request"})
	}

	if err := h.service.Create(&category); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"data":   category,
	})
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	existing, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}

	var req models.Category
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request"})
	}

	existing.Name = req.Name

	if err := h.service.Update(existing); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": existing})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found"})
	}

	if err := h.service.Delete(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Category deleted",
	})
}
