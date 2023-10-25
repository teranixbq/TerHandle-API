package entity

import (
	"terhandle/internal/app/model"
)

func UserCoreToUserModel(dataCore Core) model.Users {
	return model.Users{
		Role:         dataCore.Role,
		Profesi:      dataCore.Profesi,
		Nama:         dataCore.Nama,
		NIK:          dataCore.NIK,
		Alamat:       dataCore.Alamat,
		Longitude:    dataCore.Longitude,
		Latitude:     dataCore.Latitude,
		No_telp:      dataCore.No_telp,
		Email:        dataCore.Email,
		Password:     dataCore.Password,
		Achievements: ListAchivementCoreToModelAchivement(dataCore.Achievements),
		Status:       dataCore.Status,
		CreatedAt:    dataCore.CreatedAt,
		UpdatedAt:    dataCore.UpdatedAt,
	}
}

func ListAchivementCoreToModelAchivement(dataCore []AchievementCore) []model.Achievement {
	var result []model.Achievement
	for _, value := range dataCore {
		var UserAchievement = model.Achievement{
			UsersID:   value.UsersID,
			Nama:      value.Nama,
			Deskripsi: value.Deskripsi,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		result = append(result, UserAchievement)
	}
	return result
}

func UserModelToUserCore(dataModel model.Users) Core {
	return Core{
		Id:           int(dataModel.ID),
		Role:         dataModel.Role,
		Profesi:      dataModel.Profesi,
		Nama:         dataModel.Nama,
		NIK:          dataModel.NIK,
		Alamat:       dataModel.Alamat,
		Longitude:    dataModel.Longitude,
		Latitude:     dataModel.Latitude,
		No_telp:      dataModel.No_telp,
		Email:        dataModel.Email,
		Password:     dataModel.Password,
		Achievements: ListModelAchivementToAchivementCore(dataModel.Achievements),
		Status:       dataModel.Status,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}
}

func ListModelAchivementToAchivementCore(dataModel []model.Achievement) []AchievementCore {
	var result []AchievementCore
	for _, value := range dataModel {
		var UserAchievement = AchievementCore{
			Id:        value.ID,
			UsersID:   value.UsersID,
			Nama:      value.Nama,
			Deskripsi: value.Deskripsi,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		result = append(result, UserAchievement)
	}
	return result
}

func UserModelToUserCoreList(dataModel []model.Users) []Core {
	var ListUserModel []Core
	for _, v := range dataModel {
		ListUserModel = append(ListUserModel, UserModelToUserCore(v))
	}
	return ListUserModel
}
