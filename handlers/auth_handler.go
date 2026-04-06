package handlers

import (
	models "asset-management/model"
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Index(c *fiber.Ctx) error {
	search := c.Query("search", "")
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	users, total, err := h.service.ListUsers(search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"total":  total,
		"data":   users,
	})
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	// Kita buat struct sementara untuk menangkap password dari JSON
	type CreateRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"` // Password ditangkap di sini
		Role     string `json:"role"`
	}

	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request body"})
	}

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Kirim password mentah ke service untuk di-hash
	if err := h.service.CreateUser(&user, req.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "data": user})
}

func (h *UserHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	if err := h.service.UpdateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	if err := h.service.DeleteUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
func (h *UserHandler) Login(c *fiber.Ctx) error {
	// Struct untuk menangkap input login
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid request body"})
	}

	// Panggil service login
	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Jika berhasil, kirim data user (password akan otomatis tersembunyi karena tag json:"-")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "login successful",
		"data":    user,
	})
}
