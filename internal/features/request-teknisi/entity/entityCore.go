package entity

import (
	_feedback"terhandle/internal/features/feedback/entity"
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
	Feedback            []_feedback.CoreFeedback
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type FotoCore struct {
	Id               int
	RequestTeknisiID uint
	Foto             string
}
