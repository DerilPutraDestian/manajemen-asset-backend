package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("RahasiaNegara123")

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil header Authorization
		authHeader := c.Get("Authorization")

		// Cek apakah header ada dan formatnya benar (Bearer <token>)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Missing or invalid security token",
			})
		}

		// 2. Bersihkan string token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Parse dan Validasi Token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Pastikan algoritma yang digunakan adalah HS256
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return jwtSecret, nil
		})

		// Jika token error atau tidak valid
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Token expired or tampered",
			})
		}

		// 4. Ekstrak Claims (Data di dalam token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token payload",
			})
		}

		// 5. SIMPAN DATA KE LOCALS
		// Ini sangat penting agar Handler bisa memanggil c.Locals("role")
		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])

		// Lanjutkan ke Handler berikutnya
		return c.Next()
	}
}
