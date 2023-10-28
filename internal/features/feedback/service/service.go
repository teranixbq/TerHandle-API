package service

import (
	
	"terhandle/internal/features/feedback/entity"
	//_requestTeknisi "terhandle/internal/features/request-teknisi/entity"
)

type feedbackService struct {
	feedbackRepository entity.FeedbackRepositoryInterface
	//requestTeknisi     _requestTeknisi.UserRequestRepositoryInterface
}

func NewUserService(fr entity.FeedbackRepositoryInterface) entity.FeedbackServiceInterface {
	return &feedbackService{
		feedbackRepository: fr,
	}
}

func (fr *feedbackService) Create(payload entity.CoreFeedback) error {

	err := fr.feedbackRepository.SelectByIdAndRole(payload.UsersID, payload.TeknisiID, "user", "teknisi")
	if err != nil {
		return err
	}

	err = fr.feedbackRepository.SelectByIdAndIdUser(payload.RequestTeknisiID, payload.UsersID, payload.TeknisiID)
	if err != nil {
		return err
	}

	err = fr.feedbackRepository.Insert(payload.RequestTeknisiID, payload)
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
