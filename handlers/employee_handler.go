package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

// GET /api/employees
func (h *EmployeeHandler) Index(c *fiber.Ctx) error {
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListEmployees(search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"count":  total,
		"data":   data,
	})
}

// POST /api/employees
func (h *EmployeeHandler) Store(c *fiber.Ctx) error {
	var emp models.Employee

	if err := c.BodyParser(&emp); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}

	if err := h.service.CreateEmployee(&emp); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": emp})
}
