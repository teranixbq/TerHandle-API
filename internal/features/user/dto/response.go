package dto

import (
	"terhandle/internal/features/user/entity"
	//_feedback"terhandle/internal/features/feedback/entity"
	_feedbackdto "terhandle/internal/features/feedback/dto"
)

type ResponseLogin struct {
	Name string `json:"name"`
}

type ResponGetAll struct {
	Id       int
	Profesi  string
	Nama     string
	Alamat   string
	Email    string
	Status   string
	Feedback []_feedbackdto.ResponseFeedback
}

func CoreToResponse(dataCore entity.Core) ResponGetAll {
	return ResponGetAll{
		Id:       dataCore.Id,
		Profesi:  dataCore.Profesi,
		Nama:     dataCore.Nama,
		Alamat:   dataCore.Alamat,
		Email:    dataCore.Email,
		Status:   dataCore.Status,
		Feedback: _feedbackdto.CoreToResponseFeedbackList(dataCore.Feedback),
	}
}

// func ListFeedbackCoreToResponList(dataCore []_feedback.CoreFeedback) []_feedbackdto.ResponseFeedback {
// 	var result []_feedbackdto.ResponseFeedback
// 	for _, value := range dataCore {
// 		var UserFeedback = _feedbackdto.ResponseFeedback{

// 			RequestTeknisiID: value.RequestTeknisiID,
// 			UsersID: value.UsersID,
// 			TeknisiID: value.TeknisiID,
// 			Rating: value.Rating,
// 			Ulasan: value.Ulasan,

// 		}
// 		result = append(result, UserFeedback)
// 	}
// 	return result
// }

func CoreToResponseList(dataCore []entity.Core) []ResponGetAll {
	var result []ResponGetAll
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}

type ResponGetById struct {
	Id      int
	Profesi string
	Nama    string
	Alamat  string
	Email   string
	Status  string
}

func CoreToResponseById(dataCore entity.Core) ResponGetById {
	return ResponGetById{
		Id:      dataCore.Id,
		Profesi: dataCore.Profesi,
		Nama:    dataCore.Nama,
		Alamat:  dataCore.Alamat,
		Email:   dataCore.Email,
		Status:  dataCore.Status,
	}
}

func CoreToResponseByIdList(dataCore []entity.Core) []ResponGetById{
	var result []ResponGetById
	for _, v := range dataCore {
		result = append(result, CoreToResponseById(v))
	}
	return result
}
