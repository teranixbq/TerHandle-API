package service

import (
	"errors"
	"terhandle/internal/features/feedback/entity"
)

type feedbackService struct {
	feedbackRepository entity.FeedbackRepositoryInterface
}

func NewUserService(fr entity.FeedbackRepositoryInterface) entity.FeedbackServiceInterface {
	return &feedbackService{
		feedbackRepository: fr,
	}
}

func (fr *feedbackService) Create(payload entity.CoreFeedback) error {

	if payload.TeknisiID == 0 || payload.UsersID == 0 || payload.RequestTeknisiID == 0 || payload.Ulasan == "" {
		return errors.New("harap lengkapi data dengan benar")
	}

	if payload.Rating > 5 {
		return errors.New("tidak boleh lebih dari lima")
	}
	
	err := fr.feedbackRepository.SelectByIdAndRole(payload.UsersID, payload.TeknisiID, "user", "teknisi")
	if err != nil {
		return err
	}

	err = fr.feedbackRepository.SelectByIdAndIdUser(payload.RequestTeknisiID, payload.UsersID, payload.TeknisiID)
	if err != nil {
		return err
	}

	err = fr.feedbackRepository.Insert(payload.RequestTeknisiID,payload.TeknisiID,payload)
	if err != nil {
		return err
	}

	return nil
}

func (fr *feedbackService) UpdateFeedback(userid ,id_feedback,id_request uint, payload entity.CoreFeedback) error{

	err := fr.feedbackRepository.UpdateFeedback(userid,id_feedback,id_request,payload)
	if err != nil {
		return err
	}

	return nil
}
