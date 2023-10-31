package dto

import (
	_feedbackdto "terhandle/internal/features/feedback/dto"
	"terhandle/internal/features/user/entity"
)

type ResponseLogin struct {
	Name string `json:"name"`
}

type ResponGetAll struct {
	Id       int                             `json:"id"`
	Rating   float64                         `json:"rating"`
	Profesi  string                          `json:"profesi"`
	Nama     string                          `json:"nama"`
	Alamat   string                          `json:"alamat"`
	Email    string                          `json:"email"`
	Status   string                          `json:"status"`
	Feedback []_feedbackdto.ResponseFeedback `json:"feedback"`
}

func CoreToResponse(dataCore entity.Core) ResponGetAll {
	return ResponGetAll{
		Id:       dataCore.Id,
		Rating:   dataCore.Rating,
		Profesi:  dataCore.Profesi,
		Nama:     dataCore.Nama,
		Alamat:   dataCore.Alamat,
		Email:    dataCore.Email,
		Status:   dataCore.Status,
		Feedback: _feedbackdto.CoreToResponseFeedbackList(dataCore.Feedback),
	}
}

func CoreToResponseList(dataCore []entity.Core) []ResponGetAll {
	var result []ResponGetAll
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}

type ResponGetById struct {
	Id      int     `json:"id"`
	Rating  float64 `json:"rating"`
	Profesi string  `json:"profesi"`
	Nama    string  `json:"nama"`
	Alamat  string  `json:"alamat"`
	Email   string  `json:"email"`
	Status  string  `json:"status"`
}

func CoreToResponseById(dataCore entity.Core) ResponGetById {
	return ResponGetById{
		Id:      dataCore.Id,
		Rating:  dataCore.Rating,
		Profesi: dataCore.Profesi,
		Nama:    dataCore.Nama,
		Alamat:  dataCore.Alamat,
		Email:   dataCore.Email,
		Status:  dataCore.Status,
	}
}

func CoreToResponseByIdList(dataCore []entity.Core) []ResponGetById {
	var result []ResponGetById
	for _, v := range dataCore {
		result = append(result, CoreToResponseById(v))
	}
	return result
}

// Response Profile
type ResponProfileUser struct {
	Id        uint    `json:"id"`
	Nama      string  `json:"nama"`
	NIK       int     `json:"nik"`
	Alamat    string  `json:"alamat"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	No_telp   int     `json:"no_telp"`
	Email     string  `json:"email"`
}

func CoreToResponseProfileUser(dataCore entity.Core) ResponProfileUser {
	return ResponProfileUser{
		Id:        uint(dataCore.Id),
		Nama:      dataCore.Nama,
		NIK:       dataCore.NIK,
		Alamat:    dataCore.Alamat,
		Longitude: dataCore.Longitude,
		Latitude:  dataCore.Latitude,
		No_telp:   dataCore.No_telp,
		Email:     dataCore.Email,
	}
}

type ResponProfileTeknisi struct {
	Id        uint    `json:"id"`
	Rating    float64 `json:"rating"`
	Profesi   string  `json:"profesi"`
	Nama      string  `json:"nama"`
	NIK       int     `json:"nik"`
	Alamat    string  `json:"alamat"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	No_telp   int     `json:"no_telp"`
	Email     string  `json:"email"`
	Status    string  `json:"status"`
}

func CoreToResponseProfileTeknisi(dataCore entity.Core) ResponProfileTeknisi {
	return ResponProfileTeknisi{
		Id:        uint(dataCore.Id),
		Rating:    dataCore.Rating,
		Profesi:   dataCore.Profesi,
		Nama:      dataCore.Nama,
		NIK:       dataCore.NIK,
		Alamat:    dataCore.Alamat,
		Longitude: dataCore.Longitude,
		Latitude:  dataCore.Latitude,
		No_telp:   dataCore.No_telp,
		Email:     dataCore.Email,
		Status:    dataCore.Status,
	}
}

func ResponsesProfileList(data entity.Core, role string) interface{} {
	if role == "user" {
		return CoreToResponseProfileUser(data)
	}
	if role == "teknisi" {
		return CoreToResponseProfileTeknisi(data)
	}
	return nil
}
