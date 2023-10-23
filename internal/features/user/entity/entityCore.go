package entity

import (
	"time"
)

type Core struct {
	Id          int
	Role        string
	Name        string `validate:"required"`
	NIK         int
	Address     string
	No_telp     int
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Achievement []AchievementCore
	Status		string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AchievementCore struct {
	Id        uint
	Name      string
	Deskripsi string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
