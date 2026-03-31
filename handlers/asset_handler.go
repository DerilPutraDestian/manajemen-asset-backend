package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssetHandler struct {
	service service.AssetService
}

func NewAssetHandler(s service.AssetService) *AssetHandler {
	return &AssetHandler{s}
}

func (h *AssetHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	data, total, _ := h.service.GetAll(page, limit)

	return c.JSON(fiber.Map{
		"data":  data,
		"total": total,
	})
}

func (h *AssetHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, _ := h.service.GetByID(id)
	return c.JSON(data)
}

func (h *AssetHandler) Create(c *fiber.Ctx) error {
	var req models.CreateAssetRequest
	c.BodyParser(&req)

	if err := utils.Validate.Struct(req); err != nil {
		return c.JSON(err.Error())
	}

	asset := models.Asset{
		Code:       req.Code,
		Name:       req.Name,
		Status:     req.Status,
		Condition:  req.Condition,
		CategoryID: req.CategoryID,
	}

	h.service.Create(&asset)
	return c.JSON(asset)
}

func (h *AssetHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req models.CreateAssetRequest
	c.BodyParser(&req)

	asset := models.Asset{
		Code:       req.Code,
		Name:       req.Name,
		Status:     req.Status,
		Condition:  req.Condition,
		CategoryID: req.CategoryID,
	}

	h.service.Update(id, &asset)
	return c.JSON(asset)
}

func (h *AssetHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	h.service.Delete(id)
	return c.JSON(fiber.Map{"message": "deleted"})
}
