package service

import (
	"testing"

	"terhandle/internal/features/admin/entity"
	test "terhandle/mocks"
	"github.com/golang/mock/gomock"
)

func TestAdminService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)

	payload := entity.CoreProfesi{}

	mockRepo.EXPECT().Insert(payload).Return(nil)

	err := adminService.Create(payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}


func TestAdminService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)


	mockRepo.EXPECT().SelectAll().Return([]entity.CoreProfesi{}, nil)

	_, err := adminService.GetAll()
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}

}

func TestAdminService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)


	idProfesi := uint(1) 
	payload := entity.CoreProfesi{} 

	mockRepo.EXPECT().Update(idProfesi, payload).Return(nil)

	err := adminService.Update(idProfesi, payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}

func TestAdminService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)


	idProfesi := uint(1)

	mockRepo.EXPECT().Delete(idProfesi).Return(nil)

	err := adminService.Delete(idProfesi)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}

func TestAdminService_InsertBiaya(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)


	payload := entity.CoreTransport{}

	mockRepo.EXPECT().InsertBiaya(payload).Return(nil)

	err := adminService.InsertBiaya(payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}

func TestAdminService_UpdateBiaya(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockAdminRepositoryInterface(ctrl)
	adminService := NewUserService(mockRepo)

	id := uint(1) 
	payload := entity.CoreTransport{} 
	mockRepo.EXPECT().UpdateBiaya(id, payload).Return(nil)

	err := adminService.UpdateBiaya(id, payload)
	if err != nil {
		t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
	}
}
