package handler

import (
	"os"
	"strconv"

	"latihan/config"
	"latihan/model"
	"latihan/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Format data salah"})
	}

	user, err := repository.FindUserByUsername(req.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	// compare password (req) vs hash (db)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	expMin, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_MINUTES"))
	token, err := config.GenerateToken(user.ID, user.Username, user.Role, expMin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat token", "error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
		"user": fiber.Map{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
