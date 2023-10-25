package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Role         string `gorm:"type:enum('admin', 'user', 'teknisi');default:'user'"`
	Profesi      string `json:"profesi" form:"profesi"`
	Nama         string `json:"nama" form:"nama"`
	NIK          int    `json:"nik" form:"nik"`
	Alamat       string `json:"alamat" form:"alamat"`
	Longitude    string `json:"longitude" form:"longitude"`
	Latitude     string `json:"latitude" form:"latitude"`
	No_telp      int    `json:"no_telp" form:"no_telp"`
	Email        string `json:"email" form:"email" `
	Password     string `json:"password" form:"password"`
	Achievements []Achievement
	Status       string    `gorm:"type:enum('verified','unverified','unverified-teknisi');default:'unverified'"`
	CreatedAt    time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt    time.Time `gorm:"type:DATETIME(0)"`
}

type Achievement struct {
	gorm.Model
	UsersID   uint `json:"user_id" form:"user_id"`
	Nama      string
	Deskripsi string
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}

type RequestTeknisi struct {
	gorm.Model
	UsersID       uint `json:"user_id" form:"user_id" gorm:"foreignKey:ID"`
	TeknisiID     uint `json:"teknisi_id" form:"teknisi_id" gorm:"foreignKey:ID"`
	Foto          []Foto
	Deskripsi     string `json:"deskripsi" form:"deskripsi"`
	Jarak         int
	Biaya         float64
	Status        string    `gorm:"type:enum('menunggu diproses','diproses', 'ditolak', 'selesai');default:'menunggu diproses'"`
	CreatedAt     time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt     time.Time `gorm:"type:DATETIME(0)"`
	RequesterUser Users     `gorm:"foreignKey:UsersID"`
	TargetUser    Users     `gorm:"foreignKey:TeknisiID"`
}

type Foto struct {
	gorm.Model
	RequestTeknisiID uint
	Foto             string
}

type Voucher struct {
	gorm.Model
	Name      string
	Deskripsi string
	Potongan  float64
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}

// func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
//     currentTime := time.Now()
//     // Ubah format waktu sesuai dengan "dd-MM-yy 15:04:05"
//     formattedTime := currentTime.Format("02-01-06 15:04:05")
//     u.CreatedAt = currentTime
//     return
// }
