package dto

import "terhandle/internal/features/feedback/entity"

type RequestCreate struct {
	RequestTeknisiID uint   `json:"request_id" form:"request_id"`
	UsersID          uint   `json:"user_id" form:"user_id"`
	TeknisiID        uint   `json:"teknisi_id" form:"teknisi_id"`
	Rating           int    `json:"rating" form:"rating"`
	Ulasan           string `json:"ulasan" form:"ulasan"`
}

func RequestCreateToCore(dataRequest RequestCreate) entity.CoreFeedback {
	return entity.CoreFeedback{
		RequestTeknisiID: dataRequest.RequestTeknisiID,
		UsersID:          dataRequest.UsersID,
		TeknisiID:        dataRequest.TeknisiID,
		Rating:           dataRequest.Rating,
		Ulasan:           dataRequest.Ulasan,
	}
}
