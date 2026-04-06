package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssetHandler struct {
	service *service.AssetService
}

func NewAssetHandler(s *service.AssetService) *AssetHandler {
	return &AssetHandler{service: s}
}

func (h *AssetHandler) Index(c *fiber.Ctx) error {
	categoryID := c.Query("category_id")
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	assets, count, err := h.service.ListAssets(categoryID, search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"data":   assets,
	})
}

// GET /api/assets/:id
func (h *AssetHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	asset, err := h.service.GetAsset(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found"})
	}

	return c.JSON(fiber.Map{"status": "success", "data": asset})
}

// POST /api/assets
func (h *AssetHandler) Store(c *fiber.Ctx) error {
	var asset models.Asset
	if err := c.BodyParser(&asset); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}

	// ID akan di-generate oleh BeforeCreate di model
	if err := h.service.CreateAsset(&asset); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": asset})
}

// PUT /api/assets/:id
func (h *AssetHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Cari data lama
	asset, err := h.service.GetAsset(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found"})
	}

	// 2. Timpa dengan data baru dari Body
	if err := c.BodyParser(&asset); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// 3. Simpan perubahan
	if err := h.service.UpdateAsset(asset); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": asset})
}

// DELETE /api/assets/:id
func (h *AssetHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	asset, err := h.service.GetAsset(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found"})
	}

	if err := h.service.DeleteAsset(asset); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Asset deleted"})
}
