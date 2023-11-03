package service

import (
	"errors"

	"terhandle/internal/features/user/entity"
	test "terhandle/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockRepo)

	payload := entity.Core{
		Email:    "test@example.com",
		Password: "password",
		Nama:     "Test User",
	}

	mockRepo.EXPECT().Insert(payload).Return(nil)

	err := userService.Create(payload)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestUserService_Create_DuplicateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockRepo)

	payload := entity.Core{
		Email:    "test@example.com",
		Password: "password",
		Nama:     "Test User",
	}

	mockRepo.EXPECT().Insert(payload).Return(errors.New("email sudah terdaftar. Silakan gunakan email lain"))

	err := userService.Create(payload)
	expectedErr := "email sudah terdaftar. Silakan gunakan email lain"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("Expected error: %s, but got: %v", expectedErr, err)
	}
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockRepo)

	email := "test@example.com"
	password := "password"

	expectedUser := entity.Core{}
	expectedToken := "testToken"
	expectedErr := errors.New("error saat login")

	mockRepo.EXPECT().Login(email, password).Return(expectedUser, expectedToken, expectedErr)

	_, _, err := userService.Login(email, password)
	if err != expectedErr {
		t.Errorf("Seharusnya mendapat error %v, tetapi mendapat %v", expectedErr, err)
	}

}

func TestUserService_Login_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockUserRepositoryInterface(ctrl)
	us := NewUserService(mockRepo)

	email := ""
	password := ""

	_, _, err := us.Login(email, password)
	expectedError := errors.New("email dan password harus diisi")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}

func TestUserService_CreateUpdateDetail_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := uint(1)
    userData := entity.Core{} // Ganti dengan data yang sesuai

    expectedError := errors.New("error saat pembaruan detail user")
    mockRepo.EXPECT().Update(userID, userData).Return(expectedError)

    err := userService.CreateUpdateDetail(userID, userData)
    if err == nil || err != expectedError {
        t.Errorf("Seharusnya mendapat error: %v, tetapi mendapat: %v", expectedError, err)
    }
}

func TestUserService_GetAll(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    expectedUsers := []entity.Core{
    }

    mockRepo.EXPECT().SelectAll().Return(expectedUsers, nil)

    users, err := userService.GetAll()
    if err != nil {
        t.Errorf("Seharusnya tidak mendapat error, tapi mendapat: %v", err)
    }

    if len(users) != len(expectedUsers) {
        t.Errorf("Jumlah pengguna yang diterima tidak cocok. Got: %d, Expected: %d", len(users), len(expectedUsers))
    }
}

func TestUserService_GetAll_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    expectedError := errors.New("error saat mengambil data")

    mockRepo.EXPECT().SelectAll().Return(nil, expectedError)

    _, err := userService.GetAll()
    if err == nil || err != expectedError {
        t.Errorf("Seharusnya mendapat error: %v, tetapi mendapat: %v", expectedError, err)
    }
}


func TestUserService_GetUserById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := test.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockRepo)

	userID := uint(1)
	role := "admin"

	mockRepo.EXPECT().SelectUserById(userID, role).Return(entity.Core{}, errors.New("error select user by ID"))

	_, err := userService.GetUserById(userID, role)
	if err == nil {
		t.Error("Seharusnya mendapat error, tetapi tidak ada")
	}
}

func TestUserService_GetById(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := uint(1) 

    expectedUsers := []entity.Core{} 

    mockRepo.EXPECT().SelectByIdWithFeedback(userID).Return(expectedUsers, nil)

    _, err := userService.GetById(userID)
    if err != nil {
        t.Errorf("Seharusnya tidak mendapat error, tetapi mendapat: %v", err)
    }

}

func TestUserService_GetById_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := uint(1) 

    expectedError := errors.New("error saat memperoleh data pengguna berdasarkan ID")

    mockRepo.EXPECT().SelectByIdWithFeedback(userID).Return(nil, expectedError)

    users, err := userService.GetById(userID)
    if err == nil {
        t.Error("Seharusnya mendapat error, tetapi tidak ada")
    }

    if users != nil {
        t.Error("Users seharusnya nil ketika error terjadi")
    }
}

func TestUserService_RequestTeknisiRole(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := uint(1)

    mockRepo.EXPECT().UpdateField(userID, "role", "teknisi").Return(nil)

    err := userService.RequestTeknisiRole(userID)
    if err != nil {
        t.Errorf("Seharusnya tidak mendapat error, tetapi mendapat: %v", err)
    }
}

func TestUserService_DeleteById(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := test.NewMockUserRepositoryInterface(ctrl)
    userService := NewUserService(mockRepo)

    userID := uint(1)
    mockRepo.EXPECT().DeleteById(userID).Return(nil)

    err := userService.DeleteById(userID)
    if err != nil {
        t.Errorf("Seharusnya tidak mendapat error, tetapi mendapat: %v", err)
    }
}
