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

// GET /api/maintenances?asset_id=xxx
func (h *MaintenanceHandler) Index(c *fiber.Ctx) error {
	// Ambil asset_id dari query param untuk memfilter histori maintenance
	assetID := c.Query("asset_id")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	// Pastikan service ListByAsset menerima assetID sesuai revisi service sebelumnya
	data, total, err := h.service.ListByAsset(assetID, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"total":  total,
		"data":   data,
	})
}

// POST /api/maintenances
func (h *MaintenanceHandler) Store(c *fiber.Ctx) error {
	var maintenance models.Maintenance

	if err := c.BodyParser(&maintenance); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request body",
		})
	}

	if err := h.service.CreateMaintenance(&maintenance); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"data":   maintenance,
	})
}

// PUT /api/maintenances/:id
func (h *MaintenanceHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Ambil data lama dari database
	existing, err := h.service.GetMaintenance(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Maintenance not found"})
	}

	// 2. Tampung request body ke struct baru (jangan langsung ke 'existing')
	var req models.Maintenance
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request"})
	}

	// 3. Hanya update field yang memang dikirim di JSON
	// Jika di JSON tidak ada asset_id, jangan ubah asset_id yang sudah ada di DB
	if req.AssetID != "" {
		existing.AssetID = req.AssetID
	}
	if req.ReportedBy != "" {
		existing.ReportedBy = req.ReportedBy
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.Status != "" {
		existing.Status = req.Status
	}
	if req.EndDate != nil {
		existing.EndDate = req.EndDate
	}
	if req.StartDate != nil {
		existing.StartDate = req.StartDate
	}

	// 4. Simpan kembali data 'existing' yang sudah diperbarui field tertentu saja
	if err := h.service.UpdateMaintenance(existing); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": existing})
}
