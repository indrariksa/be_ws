package repository

import (
	"latihan/config"
	"latihan/model"
)

// Ambil semua data mahasiswa
func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	var data []model.Mahasiswa
	result := config.GetDB().Find(&data)
	return data, result.Error
}

// Insert mahasiswa baru
func InsertMahasiswa(mhs model.Mahasiswa) error {
	result := config.GetDB().Create(&mhs)
	return result.Error
}

// Ambil satu data mahasiswa berdasarkan NPM
func GetMahasiswaByNPM(npm string) (model.Mahasiswa, error) {
	var mhs model.Mahasiswa
	result := config.GetDB().First(&mhs, "npm = ?", npm)
	return mhs, result.Error
}

// Update data mahasiswa berdasarkan NPM
func UpdateMahasiswa(npm string, newData model.Mahasiswa) (model.Mahasiswa, error) {
	db := config.GetDB()

	var mhs model.Mahasiswa

	// Cari dulu data yang lama
	if err := db.First(&mhs, "npm = ?", npm).Error; err != nil {
		return mhs, err
	}

	// Update field-field yang boleh diubah
	mhs.Nama = newData.Nama
	mhs.Prodi = newData.Prodi
	mhs.Alamat = newData.Alamat
	mhs.Hobi = newData.Hobi

	// Simpan ke database
	if err := db.Save(&mhs).Error; err != nil {
		return mhs, err
	}

	return mhs, nil
}

// Hapus data mahasiswa berdasarkan NPM
func DeleteMahasiswa(npm string) error {
	result := config.GetDB().Where("npm = ?", npm).Delete(&model.Mahasiswa{})
	return result.Error
}
