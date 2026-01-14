package router

import (
	"latihan/config"
	"latihan/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// PUBLIC
	api.Post("/login", handler.Login)
	api.Post("/register", handler.CreateUser)

	// PROTECTED
	protected := api.Group("", config.JWTMiddleware())

	// Mahasiswa (butuh token)
	protected.Get("/mahasiswa", handler.GetAllMahasiswa)
	protected.Post("/mahasiswa", handler.InsertMahasiswa)
	protected.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	protected.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)
	protected.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)
}
