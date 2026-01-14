package main

import (
	"latihan/config"
	"latihan/model"
	"latihan/router"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi koneksi ke DB
	config.InitDB()

	// Auto migrate tabel
	config.DB.AutoMigrate(&model.Mahasiswa{})
	config.DB.AutoMigrate(&model.User{})

	// Setup CORS
	config.SetupCORS(app)

	// Logging request
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}
