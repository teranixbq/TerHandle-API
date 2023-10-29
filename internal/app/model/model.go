package model

import (
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Role      string `gorm:"type:enum('admin', 'user', 'teknisi');default:'user'"`
	Rating    float64	
	Profesi   string     `json:"profesi" form:"profesi"`
	Nama      string     `json:"nama" form:"nama"`
	NIK       int        `json:"nik" form:"nik"`
	Alamat    string     `json:"alamat" form:"alamat"`
	Longitude float64    `json:"longitude" form:"longitude"`
	Latitude  float64    `json:"latitude" form:"latitude"`
	No_telp   int        `json:"no_telp" form:"no_telp"`
	Email     string     `json:"email" form:"email" `
	Password  string     `json:"password" form:"password"`
	Status    string     `gorm:"type:enum('offline','online');default:'online'"`
	Feedback  []Feedback `gorm:"foreignKey:TeknisiID"`
	CreatedAt time.Time  `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time  `gorm:"type:DATETIME(0)"`
}

type RequestTeknisi struct {
	gorm.Model
	UsersID             uint    `json:"user_id" form:"user_id" gorm:"foreignKey:ID"`
	TeknisiID           uint    `json:"teknisi_id" form:"teknisi_id" gorm:"foreignKey:ID"`
	Foto                []Foto  `gorm:"foreignKey:RequestTeknisiID"`
	Deskripsi           string  `json:"deskripsi" form:"deskripsi"`
	Jarak               float64 `json:"jarak" form:"jarak"`
	Biaya               float64
	Status              string     `gorm:"type:enum('menunggu diproses','konfirmasi biaya','menunggu konfimasi','diproses', 'dibatalkan', 'selesai');default:'menunggu diproses'"`
	Feedback            []Feedback `gorm:"foreignKey:RequestTeknisiID"`
	Diproses            bool       `json:"diproses" form:"diproses"`
	Konfirmasi_biaya    bool       `json:"konfirmasi_biaya" form:"konfirmasi_biaya"`
	Menunggu_konfirmasi bool       `json:"menunggu_konfirmasi" form:"menunggu_konfirmasi"`
	Dibatalkan          bool       `json:"dibatalkan" form:"dibatalkan"`
	Selesai             bool       `json:"selesai" form:"selesai"`
	CreatedAt           time.Time  `gorm:"type:DATETIME(0)"`
	UpdatedAt           time.Time  `gorm:"type:DATETIME(0)"`
	RequesterUser       Users      `gorm:"foreignKey:UsersID"`
	TargetUser          Users      `gorm:"foreignKey:TeknisiID"`
}

type Foto struct {
	gorm.Model
	RequestTeknisiID uint
	Foto             string `form:"foto"`
}

type Feedback struct {
	gorm.Model
	RequestTeknisiID uint           `json:"request_id" form:"request_id" gorm:"foreignKey:ID"`
	UsersID          uint           `json:"user_id" form:"user_id" gorm:"foreignKey:ID"`
	TeknisiID        uint           `json:"teknisi_id" form:"teknisi_id" gorm:"foreignKey:ID"`
	Rating           int            `json:"rating" form:"rating" gorm:"not null"`
	Ulasan           string         `json:"ulasan" form:"ulasan"`
	CreatedAt        time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt        time.Time      `gorm:"type:DATETIME(0)"`
	FeedbackUser     Users          `gorm:"foreignKey:UsersID"`
	TargetUser       Users          `gorm:"foreignKey:TeknisiID"`
	RequestUser      RequestTeknisi `gorm:"foreignKey:RequestTeknisiID"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	myUUID := uuid.New()

	hash := sha256.Sum256(myUUID[:])

	Iduint := binary.BigEndian.Uint32(hash[:16])

	u.ID = uint(Iduint)

	return nil
}
