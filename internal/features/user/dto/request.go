package dto

import "terhandle/internal/features/user/entity"

type RequestCreate struct {
	Profesi  string `json:"profesi" form:"profesi" gorm:"-"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(dataRequest RequestCreate) entity.Core {
	return entity.Core{
		Profesi:  dataRequest.Profesi,
		Nama:     dataRequest.Nama,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}
}

type RequestCreateDetail struct {
	Profesi   string
	NIK       int
	Alamat    string
	Longitude float64
	Latitude  float64
	No_telp   int
	Email     string
	Status    string
	Password string `json:"password" form:"password"`
}

func RequestCreateDetailToCore(dataRequest RequestCreateDetail) entity.Core {
	return entity.Core{
		Profesi:   dataRequest.Profesi,
		NIK:       dataRequest.NIK,
		Alamat:    dataRequest.Alamat,
		Longitude: dataRequest.Longitude,
		Latitude:  dataRequest.Latitude,
		No_telp:   dataRequest.No_telp,
		Email:     dataRequest.Email,
		Status:    dataRequest.Status,
		Password: dataRequest.Password,
	}
}

type RequestLogin struct {
	Id       int    `json:"id" form:"id"`
	Role     string `json:"role" form:"role"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestLoginToCore(dataRequest RequestLogin) entity.Core {
	return entity.Core{
		Id:       dataRequest.Id,
		Role:     dataRequest.Role,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}
}
