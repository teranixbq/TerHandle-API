package entity

import (
	"terhandle/internal/app/model"
)

func UserCoreToUserModel(dataCore Core) model.Users {
	return model.Users{
		Role:         dataCore.Role,
		Name:         dataCore.Name,
		NIK:          dataCore.NIK,
		Address:      dataCore.Address,
		No_telp:      dataCore.No_telp,
		Email:        dataCore.Email,
		Password:     dataCore.Password,
		Achievements: ListAchivementCoreToModelAchivement(dataCore.Achievement),
		Status:			dataCore.Status,
	}
}

func ListAchivementCoreToModelAchivement(dataCore []AchievementCore) []model.Achievement {
	var result []model.Achievement
	for _, value := range dataCore {
		var UserAchievement = model.Achievement{
			Name:      value.Name,
			Deskripsi: value.Deskripsi,
		}

		result = append(result, UserAchievement)
	}
	return result
}

func UserModelToUserCore(dataModel model.Users) Core {
	return Core{
		Id:          int(dataModel.ID),
		Role:        dataModel.Role,
		Name:        dataModel.Name,
		NIK:         dataModel.NIK,
		Address:     dataModel.Address,
		No_telp:     dataModel.No_telp,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		Achievement: ListModelAchivementToAchivementCore(dataModel.Achievements),
		Status:			dataModel.Status,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

func ListModelAchivementToAchivementCore(dataModel []model.Achievement) []AchievementCore {
	var result []AchievementCore
	for _, value := range dataModel {
		var UserAchievement = AchievementCore{
			Id:        value.ID,
			Name:      value.Name,
			Deskripsi: value.Deskripsi,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}

		result = append(result, UserAchievement)
	}
	return result
}
