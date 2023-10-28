package entity

import (
	"terhandle/internal/features/feedback/entity"
	"time"
)

type Core struct {
	Id                  uint
	UsersID             uint
	TeknisiID           uint
	Foto                []FotoCore
	Deskripsi           string
	Jarak               float64
	Biaya               float64
	Diproses            bool
	Konfirmasi_biaya    bool
	Menunggu_konfirmasi bool
	Dibatalkan          bool
	Selesai             bool
	Status              string
	Feedback            []entity.CoreFeedback
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type FotoCore struct {
	Id               int
	RequestTeknisiID uint
	Foto             string
}

// type Core struct {
// 	Id          int
// 	Role         string
// 	Profesi      string
// 	Nama         string
// 	NIK          int
// 	Alamat       string
// 	Longitude    string
// 	Latitude     string
// 	No_telp      int
// 	Email        string
// 	Password     string
// 	Achievements []AchievementCore
// 	Status       string
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// }

// type AchievementCore struct {
// 	Id        uint
// 	UsersID   uint
// 	Nama      string
// 	Deskripsi string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
