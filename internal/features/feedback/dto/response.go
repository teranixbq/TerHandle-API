package dto

import "terhandle/internal/features/feedback/entity"

type ResponseFeedback struct {
	UsersID          uint
	TeknisiID        uint
	Rating           int
	Ulasan           string
}

func CoreFeedbackToResponseFeedback(dataCore entity.CoreFeedback) ResponseFeedback {
	return ResponseFeedback{
		UsersID:          dataCore.UsersID,
		TeknisiID:        dataCore.TeknisiID,
		Rating:           dataCore.Rating,
		Ulasan:           dataCore.Ulasan,
	}
}


func CoreToResponseFeedbackList(dataCore []entity.CoreFeedback) []ResponseFeedback {
	var result []ResponseFeedback
	for _, v := range dataCore {
		result = append(result, CoreFeedbackToResponseFeedback(v))
	}
	return result
}
