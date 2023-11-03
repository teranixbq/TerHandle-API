package service

import (
	"errors"
	"testing"

	"terhandle/internal/features/feedback/entity"
	test "terhandle/mocks"
	"github.com/golang/mock/gomock"

)

func TestFeedbackService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	payload := entity.CoreFeedback{
		UsersID:        1,
		TeknisiID:      2,
		RequestTeknisiID: 3,
		Ulasan:         "Ulasan",
		Rating:         4,
	} 

	mockRepo.EXPECT().SelectByIdAndRole(payload.UsersID, payload.TeknisiID, "user", "teknisi").Return(nil)
	mockRepo.EXPECT().SelectByIdAndIdUser(payload.RequestTeknisiID, payload.UsersID, payload.TeknisiID).Return(nil)
	mockRepo.EXPECT().Insert(payload.RequestTeknisiID, payload.TeknisiID, payload).Return(nil)

	err := feedbackService.Create(payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}

func TestFeedbackService_Create_ValidationErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	payload := entity.CoreFeedback{} 

	err := feedbackService.Create(payload)
	expectedError := "harap lengkapi data dengan benar"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Harapan error '%s', tetapi mendapat %v", expectedError, err)
	}
}

func TestFeedbackService_Create_RatingError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	payload := entity.CoreFeedback{
		UsersID:        1,
		TeknisiID:      2,
		RequestTeknisiID: 3,
		Ulasan:         "Ulasan",
		Rating:         6, 
	} 

	err := feedbackService.Create(payload)
	expectedError := "tidak boleh lebih dari lima"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Harapan error '%s', tetapi mendapat %v", expectedError, err)
	}
}

func TestFeedbackService_Create_SelectByIdAndRoleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	payload := entity.CoreFeedback{
		UsersID:        1,
		TeknisiID:      2,
		RequestTeknisiID: 3,
		Ulasan:         "Ulasan",
		Rating:         4,
	} 

	expectedErr := errors.New("Gagal select")

	mockRepo.EXPECT().SelectByIdAndRole(payload.UsersID, payload.TeknisiID, "user", "teknisi").Return(expectedErr)

	err := feedbackService.Create(payload)
	if err != expectedErr {
		t.Errorf("Harapan error '%v', tetapi mendapat %v", expectedErr, err)
	}
}

func TestFeedbackService_UpdateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	userID := uint(1) 
	idFeedback := uint(2) 
	idRequest := uint(3)
	payload := entity.CoreFeedback{} 

	mockRepo.EXPECT().UpdateFeedback(userID, idFeedback, idRequest, payload).Return(nil)

	err := feedbackService.UpdateFeedback(userID, idFeedback, idRequest, payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}

func TestFeedbackService_UpdateFeedback_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockFeedbackRepositoryInterface(ctrl)
	feedbackService := NewUserService(mockRepo)

	userID := uint(1) 
	idFeedback := uint(2) 
	idRequest := uint(3) 
	payload := entity.CoreFeedback{}

	expectedErr := errors.New("Gagal update")

	mockRepo.EXPECT().UpdateFeedback(userID, idFeedback, idRequest, payload).Return(expectedErr)

	err := feedbackService.UpdateFeedback(userID, idFeedback, idRequest, payload)
	if err != expectedErr {
		t.Errorf("Harapan error '%v', tetapi mendapat %v", expectedErr, err)
	}
}
