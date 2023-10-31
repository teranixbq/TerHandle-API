package entity

import (
	"terhandle/internal/app/model"
	_feedback "terhandle/internal/features/feedback/entity"
)

func UserCoreToUserModel(dataCore Core) model.Users {
	return model.Users{
		Role:      dataCore.Role,
		Rating:    dataCore.Rating,
		Profesi:   dataCore.Profesi,
		Nama:      dataCore.Nama,
		NIK:       dataCore.NIK,
		Alamat:    dataCore.Alamat,
		Longitude: dataCore.Longitude,
		Latitude:  dataCore.Latitude,
		No_telp:   dataCore.No_telp,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Status:    dataCore.Status,
		Feedback:  ListCoreFeedbackToModelFeedback(dataCore.Feedback),
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}

func ListCoreFeedbackToModelFeedback(dataCore []_feedback.CoreFeedback) []model.Feedback {
	var result []model.Feedback
	for _, value := range dataCore {
		var UserFeedback = model.Feedback{
			UsersID:   value.UsersID,
			TeknisiID: value.TeknisiID,
			Rating:    value.Rating,
			Ulasan:    value.Ulasan,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		result = append(result, UserFeedback)
	}
	return result
}

func UserModelToUserCore(dataModel model.Users) Core {
	return Core{
		Id:          int(dataModel.ID),
		Role:        dataModel.Role,
		Rating:      dataModel.Rating,
		Profesi: dataModel.Profesi,
		Nama:        dataModel.Nama,
		NIK:         dataModel.NIK,
		Alamat:      dataModel.Alamat,
		Longitude:   dataModel.Longitude,
		Latitude:    dataModel.Latitude,
		No_telp:     dataModel.No_telp,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		Feedback:    ListModelFeedbackToCoreFeedback(dataModel.Feedback),
		Status:      dataModel.Status,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

func ListModelFeedbackToCoreFeedback(dataModel []model.Feedback) []_feedback.CoreFeedback {
	var result []_feedback.CoreFeedback
	for _, value := range dataModel {
		var UserFeedback = _feedback.CoreFeedback{
			Id:        value.ID,
			UsersID:   value.UsersID,
			TeknisiID: value.TeknisiID,
			Rating:    value.Rating,
			Ulasan:    value.Ulasan,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		result = append(result, UserFeedback)
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
