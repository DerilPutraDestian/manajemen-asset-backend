package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MaintenanceHandler struct {
	service *service.MaintenanceService
}

func NewMaintenanceHandler(s *service.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{service: s}
}

func (h *MaintenanceHandler) Index(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	data, total, err := h.service.ListMaintenances(limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"total":  total,
		"data":   data,
	})
}
func (h *MaintenanceHandler) Store(c *fiber.Ctx) error {
	var m models.Maintenance
	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := h.service.CreateMaintenance(&m); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": m})
}

func (h *MaintenanceHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.service.GetDetail(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Maintenance not found"})
	}

	return c.JSON(fiber.Map{"status": "success", "data": data})
}

func (h *MaintenanceHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	m, err := h.service.GetDetail(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Maintenance not found"})
	}

	if err := c.BodyParser(m); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := h.service.UpdateMaintenance(m); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": m})
}
