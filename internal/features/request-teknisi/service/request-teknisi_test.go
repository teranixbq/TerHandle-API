package service

import (
	"errors"
	"testing"

	"terhandle/internal/features/request-teknisi/entity"
	test "terhandle/mocks"

	"github.com/golang/mock/gomock"
)

func TestUserService_Create(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    payload := entity.Core{
        TeknisiID: 1,    
        UsersID: 2,      
        Deskripsi: "Deskripsi yang valid",  
    }

    mockRepo.EXPECT().SelectByIdAndRole(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().Insert(gomock.Any()).Return(nil)

    err := userService.Create(payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_Create_RepositoryInsertError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    payload := entity.Core{
        TeknisiID: 1,
        UsersID: 2,
        Deskripsi: "Deskripsi yang valid",
    }

    mockRepo.EXPECT().SelectByIdAndRole(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().Insert(gomock.Any()).Return(errors.New("Error during insert"))

    err := userService.Create(payload)
    if err == nil {
        t.Error("Harapan error dari repository insert, tetapi tidak ada error yang diterima")
    }
}


func TestUserService_CreateUpdateDetail(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    payload := entity.Core{} 

    mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

    err := userService.CreateUpdateDetail(userID, payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

// TestUserService_GetAllById untuk pengujian GetAllById
func TestUserService_GetAllById(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 

    mockRepo.EXPECT().SelectAllById(gomock.Any()).Return([]entity.Core{}, nil)

    _, err := userService.GetAllById(userID)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_GetById(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 

    mockRepo.EXPECT().SelectById(gomock.Any(), gomock.Any()).Return([]entity.Core{}, nil)

    _, err := userService.GetById(userID, requestID)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_ClaimRequest(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    requestID := 1 
    payload := entity.Core{}

    mockRepo.EXPECT().UpdateClaims(gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

    err := userService.ClaimRequest(requestID, payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_ClaimRequest_UpdateFieldError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    requestID := 1 
    payload := entity.Core{}

    mockRepo.EXPECT().UpdateClaims(gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("Error while updating field"))

    err := userService.ClaimRequest(requestID, payload)
    if err == nil {
        t.Error("Harapan error dari UpdateField, tetapi tidak ada error yang diterima")
    }
}


func TestUserService_KonfirmasiBiaya(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusKonfirmBiaya(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

    err := userService.KonfirmasiBiaya(userID, requestID, payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_BatalkanRequest(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusBatal(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

    err := userService.BatalkanRequest(userID, requestID, payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_BatalkanRequest_UpdateStatusBatalError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusBatal(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("Error updating status to 'batal'"))

    err := userService.BatalkanRequest(userID, requestID, payload)
    if err == nil {
        t.Error("Harapan error dari UpdateStatusBatal, tetapi tidak ada error yang diterima")
    }
}

func TestUserService_BatalkanRequest_UpdateFieldError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusBatal(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), "status", "dibatalkan").Return(errors.New("Error updating field 'status' to 'dibatalkan'"))

    err := userService.BatalkanRequest(userID, requestID, payload)
    if err == nil {
        t.Error("Harapan error dari UpdateField, tetapi tidak ada error yang diterima")
    }
}


func TestUserService_SelesaikanRequest(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusSelesai(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
    mockRepo.EXPECT().UpdateField(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

    err := userService.SelesaikanRequest(userID, requestID, payload)
    if err != nil {
        t.Errorf("Harapan tidak ada error, tetapi mendapat %v", err)
    }
}

func TestUserService_SelesaikanRequestError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := 1 
    requestID := 2 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateStatusSelesai(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("Error updating status to 'selesai'"))

    err := userService.SelesaikanRequest(userID, requestID, payload)
    if err == nil {
        t.Error("Harapan error dari UpdateStatusSelesai, tetapi tidak ada error yang diterima")
    }
}

func TestUserService_Create_InvalidPayload(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    invalidPayload := entity.Core{} 

    err := userService.Create(invalidPayload)
    if err == nil {
        t.Error("Harapan error karena payload tidak valid, tetapi tidak ada error yang diterima")
    }
}

func TestUserService_ClaimRequest_UpdateClaimsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRequestRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    requestID := 1 
    payload := entity.Core{} 

    mockRepo.EXPECT().UpdateClaims(gomock.Any(), gomock.Any()).Return(errors.New("Error updating claims"))

    err := userService.ClaimRequest(requestID, payload)
    if err == nil {
        t.Error("Harapan error karena kesalahan saat update claims, tetapi tidak ada error yang diterima")
    }
}
