package service

import (
	// "errors"
	// "fmt"
	"terhandle/internal/features/request-teknisi/entity"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository entity.UserRequestRepositoryInterface
	validate       *validator.Validate
}

func NewUserService(ur entity.UserRequestRepositoryInterface) entity.UserRequestServiceInterface {

	return &userService{userRepository: ur, validate: validator.New()}
}

func (us *userService) Create(payload entity.Core) error {

	err := us.userRepository.SelectByIdAndRole(int(payload.UsersID), int(payload.TeknisiID), "user", "teknisi")
	if err != nil {
		return err
	}

	err = us.userRepository.Insert(payload)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) CreateUpdateDetail(userid int, data entity.Core) error {

	if err := us.userRepository.Update(userid, data); err != nil {
		return err
	}

	return nil
}

func (us *userService) RequestTeknisiRole(userid int) error {

	if err := us.userRepository.UpdateField(userid, "role", "teknisi"); err != nil {
		return err
	}

	return nil
}

func (us *userService) GetAllById(userid int) ([]entity.Core, error) {
	history, err := us.userRepository.SelectAllById(userid)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (us *userService) GetById(userid, id int) ([]entity.Core, error) {
	result, err := us.userRepository.SelectById(userid, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
