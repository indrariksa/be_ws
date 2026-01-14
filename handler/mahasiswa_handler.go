package handler

import (
	"latihan/model"
	"latihan/repository"

	"github.com/gofiber/fiber/v2"
)

// Get All Mahasiswa
func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil data",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data mahasiswa",
		"data":    data,
	})
}

// Insert Mahasiswa
func InsertMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	if err := repository.InsertMahasiswa(mhs); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menambah data mahasiswa",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Mahasiswa berhasil ditambahkan",
		"data":    mhs,
	})
}

// Get satu Mahasiswa berdasarkan NPM
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm") // ambil nilai :npm dari URL

	data, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Mahasiswa tidak ditemukan",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data mahasiswa",
		"data":    data,
	})
}

// Update data Mahasiswa berdasarkan NPM
func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")

	var input model.Mahasiswa
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	updated, err := repository.UpdateMahasiswa(npm, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengupdate data mahasiswa",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Mahasiswa berhasil diupdate",
		"data":    updated,
	})
}

// Delete Mahasiswa berdasarkan NPM
func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")

	if err := repository.DeleteMahasiswa(npm); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus data mahasiswa",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Mahasiswa berhasil dihapus",
	})
}
