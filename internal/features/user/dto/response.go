package dto

import "terhandle/internal/features/user/entity"

type ResponseLogin struct {
	Name string `json:"name"`
}

type ResponGetAll struct {
	Id      int
	Profesi string
	Nama    string
	Alamat  string
	// Longitude string
	// Latitude  string
	Email       string
	Achivements []Achievement
	Status      string
}

type Achievement struct {
	Nama      string
	Deskripsi string
}

func CoreToResponse(dataCore entity.Core) ResponGetAll {
	return ResponGetAll{
		Id:      dataCore.Id,
		Profesi: dataCore.Profesi,
		Nama:    dataCore.Nama,
		Alamat:  dataCore.Alamat,
		// Longitude: dataCore.Longitude,
		// Latitude:  dataCore.Latitude,
		Email:       dataCore.Email,
		Achivements: ListAchivementCoreToResponList(dataCore.Achievements),
		Status:      dataCore.Status,
	}
}

func ListAchivementCoreToResponList(dataCore []entity.AchievementCore) []Achievement {
	var result []Achievement
	for _, value := range dataCore {
		var UserAchievement = Achievement{
			Nama:      value.Nama,
			Deskripsi: value.Deskripsi,
		}
		result = append(result, UserAchievement)
	}
	return result
}

func CoreToResponseList(dataCore []entity.Core) []ResponGetAll {
	var result []ResponGetAll
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
