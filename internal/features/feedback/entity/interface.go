package entity

type FeedbackRepositoryInterface interface {
	Insert(request_id uint, data CoreFeedback) error
	SelectByIdAndIdUser(id, user_id, teknisi_id uint) error
	SelectByIdAndRole(user_id, teknisi_id uint, role_user, role_teknisi string) error
	UpdateFeedback(id_user, id_feedback, id_request uint, data CoreFeedback) error
}

type FeedbackServiceInterface interface {
	Create(data CoreFeedback) error
	UpdateFeedback(userid, id_feedback, id_request uint, data CoreFeedback) error
}
