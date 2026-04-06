package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	service *service.AssetLoanService
}

func NewLoanHandler(s *service.AssetLoanService) *LoanHandler {
	return &LoanHandler{service: s}
}

func (h *LoanHandler) Index(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := h.service.ListLoans("", "", limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"total":  total,
		"data":   data,
	})
}

func (h *LoanHandler) Store(c *fiber.Ctx) error {
	var loan models.AssetLoan

	if err := c.BodyParser(&loan); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request"})
	}

	if err := h.service.CreateLoan(&loan); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"data":   loan,
	})
}

func (h *LoanHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	existing, err := h.service.GetLoanByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Loan not found"})
	}

	var req models.AssetLoan
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request"})
	}

	existing.AssetID = req.AssetID
	existing.UserID = req.UserID
	existing.LoanDate = req.LoanDate
	existing.ReturnDate = req.ReturnDate
	existing.Status = req.Status

	if err := h.service.UpdateLoan(existing); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": existing})
}
