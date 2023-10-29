package entity

import (
	_feedback "terhandle/internal/features/feedback/entity"
	"time"
)

type Core struct {
	Id        int
	Role      string
	Profesi   string
	Nama      string
	NIK       int
	Alamat    string
	Longitude float64
	Latitude  float64
	No_telp   int
	Email     string
	Password  string
	Status    string
	Feedback  []_feedback.CoreFeedback
	CreatedAt time.Time
	UpdatedAt time.Time
}
