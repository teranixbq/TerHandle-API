package entity

import (
	"time"
)


type Core struct {
	Id          int
	Role         string 
	Profesi      string
	Nama         string 
	NIK          int   
	Alamat       string
	Longitude    string 
	Latitude     string 
	No_telp      int   
	Email        string
	Password     string 
	Achievements []AchievementCore
	Status       string    
	CreatedAt    time.Time 
	UpdatedAt    time.Time 
}


type AchievementCore struct {
	Id        uint
	UsersID   uint
	Nama      string
	Deskripsi string
	CreatedAt time.Time 
	UpdatedAt time.Time 
}