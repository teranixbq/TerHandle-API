package dto

import "terhandle/internal/features/user/entity"

type RequestCreate struct {
	Name     string `json:"name" form:"email"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(dataRequest RequestCreate) entity.Core {
	return entity.Core{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}
}

type RequestCreateDetail struct {
	Name        string
	NIK         int
	Address     string
	No_telp     int
	Email       string
	Password    string
	Achievement []entity.AchievementCore
}

func RequestCreateDetailToCore(dataRequest RequestCreateDetail) entity.Core {
	return entity.Core{
		Name:        dataRequest.Name,
		NIK:         dataRequest.NIK,
		Address:     dataRequest.Address,
		No_telp:     dataRequest.No_telp,
		Email:       dataRequest.Email,
		Password:    dataRequest.Password,
		Achievement: dataRequest.Achievement,
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
		Id:    dataRequest.Id,
		Role: dataRequest.Role,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}
}
