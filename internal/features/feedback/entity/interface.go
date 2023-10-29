package entity


type FeedbackRepositoryInterface interface {
	Insert(id_request ,id_teknisi uint, data CoreFeedback) error
	SelectByIdAndIdUser(id, user_id, teknisi_id uint) error
	SelectByIdAndRole(user_id, teknisi_id uint, role_user, role_teknisi string) error
	UpdateFeedback(id_user, id_feedback, id_request uint, data CoreFeedback) error
	//GetFeedbacksForUser(id_teknisi uint) ([]CoreFeedback, error)
}

type FeedbackServiceInterface interface {
	Create(data CoreFeedback) error
	UpdateFeedback(userid, id_feedback, id_request uint, data CoreFeedback) error
}
