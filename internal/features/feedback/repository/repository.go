package repository

import (
	"errors"
	"fmt"
	//"fmt"
	"time"

	"terhandle/internal/app/model"
	"terhandle/internal/features/feedback/entity"
	"terhandle/internal/utils/helper"

	//"terhandle/internal/utils/helper"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.FeedbackRepositoryInterface {
	return &FeedbackRepository{
		db: db,
	}
}

func (ur *FeedbackRepository) Insert(id_request ,id_teknisi uint, data entity.CoreFeedback) error {
	userRequest := model.Feedback{}
	request := model.RequestTeknisi{}
	users := model.Users{}

	if err := ur.db.Where("id = ? AND selesai = 1", id_request).First(&request).Error; err != nil {
		return err
	}

	if err := ur.db.Where("request_teknisi_id = ?", id_request).First(&userRequest).Error; err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("gagal memeriksa data: " + err.Error())
		}

		user := entity.CoreFeedbackToModelFeedback(data)
		if err := ur.db.Create(&user).Error; err != nil {
			return err
		}

		var allFeedbacks []model.Feedback
		if err := ur.db.Where("teknisi_id = ?", id_teknisi).Find(&allFeedbacks).Error; err != nil {
			return err
		}

		averageRating := helper.GetAverageRating(allFeedbacks)
		fmt.Println(users.Rating)
		if err := ur.db.Model(&users).Where("id = ?", id_teknisi).Update("rating", averageRating).Error; err != nil {
			return errors.New("gagal memperbarui rating user: " + err.Error())
		}

		return nil
	}
	return errors.New("feedback sudah ada")
}

func (ur *FeedbackRepository) SelectByIdAndIdUser(id, user_id, teknisi_id uint) error {
	userRequest := model.RequestTeknisi{}
	err := ur.db.Where("id = ? AND users_id = ? AND teknisi_id = ? ", id, user_id, teknisi_id).First(&userRequest)
	if err != nil {
		return err.Error
	}
	return nil
}

func (ur *FeedbackRepository) SelectByIdAndRole(userid, teknisiid uint, role_user, role_teknisi string) error {
	userUser := model.Users{}
	userTeknisi := model.Users{}
	roleUser := ur.db.Where("id = ? AND role = ?", userid, role_user).First(&userUser)
	if roleUser.Error != nil {
		return roleUser.Error
	}

	roleTeknisi := ur.db.Where("id = ? AND role = ?", teknisiid, role_teknisi).First(&userTeknisi)
	if roleTeknisi.Error != nil {
		return roleTeknisi.Error
	}

	return nil
}

func (ur *FeedbackRepository) UpdateFeedback(id_user, id_feedback, id_request uint, data entity.CoreFeedback) error {
	feedback := model.Feedback{}
	user := model.Users{}
	request := model.RequestTeknisi{}
	if err := ur.db.Where("id = ?", id_user).First(&user).Error; err != nil {
		return err
	}
	if err := ur.db.Where("id = ? AND selesai = 1", id_request).First(&request).Error; err != nil {
		return err
	}

	if err := ur.db.First(&feedback, id_feedback).Error; err != nil {
		return err
	}

	if time.Since(feedback.UpdatedAt) > time.Minute*5 {
		return errors.New("feedback cannot be updated as it has been more than 5 minutes")
	}

	if err := ur.db.Model(&feedback).Updates(data).Error; err != nil {
		return err
	}

	feedback.UpdatedAt = time.Now()
	if err := ur.db.Save(&feedback).Error; err != nil {
		return err
	}

	return nil
}

// func (ur *FeedbackRepository) GetFeedbacksForUser(id_teknisi uint) ([]entity.CoreFeedback, error) {
// 	var feedbacks []model.Feedback
// 	err := ur.db.Where("id = ? ", id_teknisi).Find(&feedbacks).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	feedbacksModel := entity.FeedbackModelToCoreFeedbackList(feedbacks)
// 	return feedbacksModel, nil
// }
