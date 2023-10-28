package entity

import (
	"terhandle/internal/app/model"
)

func CoreFeedbackToModelFeedback(dataCore CoreFeedback) model.Feedback {
	return model.Feedback{
		RequestTeknisiID: dataCore.RequestTeknisiID,
		UsersID:          dataCore.UsersID,
		TeknisiID:        dataCore.TeknisiID,
		Rating:           dataCore.Rating,
		Ulasan:           dataCore.Ulasan,
		CreatedAt:        dataCore.CreatedAt,
		UpdatedAt:        dataCore.UpdatedAt,
	}
}

func FeedbackModelToCoreFeedback(dataModel model.Feedback) CoreFeedback {
	return CoreFeedback{
		Id:        dataModel.ID,
		UsersID:   dataModel.UsersID,
		TeknisiID: dataModel.TeknisiID,
		Rating:    dataModel.Rating,
		Ulasan:    dataModel.Ulasan,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}


func UserModelToUserCoreList(dataModel []model.Feedback) []CoreFeedback {
	var ListFeedbackModel []CoreFeedback
	for _, v := range dataModel {
		ListFeedbackModel = append(ListFeedbackModel, FeedbackModelToCoreFeedback(v))
	}
	return ListFeedbackModel
}
