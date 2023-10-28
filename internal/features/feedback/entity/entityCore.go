package entity

import (
	"time"
)

type CoreFeedback struct {
	Id               uint
	RequestTeknisiID uint
	UsersID          uint           
	TeknisiID        uint           
	Rating           int       
	Ulasan           string          
	CreatedAt        time.Time       
	UpdatedAt        time.Time       
	
}

