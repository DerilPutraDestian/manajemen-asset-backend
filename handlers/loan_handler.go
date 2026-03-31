package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	service service.LoanService
}

func NewLoanHandler(s service.LoanService) *LoanHandler {
	return &LoanHandler{s}
}

// 🔥 BORROW
func (h *LoanHandler) Borrow(c *fiber.Ctx) error {
	var req models.LoanRequest

	// parse body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// validate
	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// call service
	if err := h.service.Borrow(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "asset borrowed successfully",
	})
}

// 🔥 RETURN
func (h *LoanHandler) Return(c *fiber.Ctx) error {
	assetID, err := strconv.Atoi(c.Params("asset_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid asset id",
		})
	}

	if err := h.service.Return(assetID); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "asset returned successfully",
	})
}
