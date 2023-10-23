package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Role         string        `gorm:"type:enum('admin', 'user', 'teknisi');default:'user'"`
	Name         string        `json:"name" form:"name"`
	NIK          int           `json:"nik" form:"nik"`
	Address      string        `json:"address" form:"address"`
	No_telp      int           `json:"no_telp" form:"no_telp"`
	Email        string        `json:"email" form:"email" gorm:"unique"`
	Password     string        `json:"password" form:"password"`
	Achievements []Achievement `gorm:"foreignKey:UserID"`
	Status       string        `gorm:"type:enum('verified','unverified','unverified-teknisi');default:'unverified'"`
}

type Admin struct {
	ID       int
	Nama     string
	Password string
	Email    string
}

type RequestTeknisi struct {
	gorm.Model
	UserID         Users `gorm:"foreignKey:UserID;references:ID"`
	TeknisiID      Users `gorm:"foreignKey:TeknisiID;references:ID"`
	Lokasi         string
	Deskripsi      DeskripsiPermasalahan
	KontakUser     string
	Biaya          float64
	Status         string
	TanggalRequest time.Time
}

type DeskripsiPermasalahan struct {
	gorm.Model
	RequestID uint
	Foto      []string
	Teks      string
}

type Achievement struct {
	gorm.Model
	UserID    uint
	Name      string
	Deskripsi string
}

type Voucher struct {
	gorm.Model
	Name      string
	Deskripsi string
	Potongan  float64
}
