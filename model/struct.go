package model

import "github.com/lib/pq"

type Mahasiswa struct {
	NPM    string         `json:"npm"    gorm:"column:npm;primaryKey;type:varchar(20);not null"`
	Nama   string         `json:"nama"   gorm:"column:nama;type:varchar(100);not null"`
	Prodi  string         `json:"prodi"  gorm:"column:prodi;type:varchar(100);not null"`
	Alamat string         `json:"alamat" gorm:"column:alamat;type:varchar(200)"`
	Hobi   pq.StringArray `json:"hobi"   gorm:"column:hobi;type:text[]"`
}

func (Mahasiswa) TableName() string { return "mahasiswa" }

type User struct {
	ID       string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"-" gorm:"column:password;not null"`
	Role     string `json:"role" gorm:"column:role;type:varchar(30);default:user"`
}

func (User) TableName() string { return "users" }

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
