package service

import (
	"errors"
	"terhandle/internal/features/user/dto"
	"terhandle/internal/features/user/entity"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

type userService struct {
	userRepository entity.UserRepositoryInterface
	validate       *validator.Validate
}

func NewUserService(ur entity.UserRepositoryInterface) entity.UserServiceInterface {

	return &userService{userRepository: ur, validate: validator.New()}
}

func (us *userService) Create(payload entity.Core) error {
	if payload.Email == "" || payload.Password == "" || payload.Name == "" {
		return errors.New("error. email dan password harus diisi")
	}

	// errValidate := us.validate.Struct(payload)

	// if errValidate != nil {
	// 	return errValidate
	// }

	err := us.userRepository.Insert(payload)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New("email sudah terdaftar. Silakan gunakan email lain")
			}
		}
		return err
	}
	return nil
}

func (us *userService) Login(email, password string) (entity.Core, string, error) {
	if email == "" || password == ""{
		return entity.Core{}, "",errors.New("email dan password harus diisi")
	}

	user := dto.RequestLogin{
		Email:    email,
		Password: password,
	}

	data, token, err := us.userRepository.Login(user.Email, user.Password)
	if err != nil {
		return data, "", err
	}
	return data, token, nil
}

func (us *userService) CreateUpdateDetail(userid int, data entity.Core) error {

	if err := us.userRepository.Update(userid, data); err != nil {
		return err
	}

	return nil
}

func (us *userService) RequestTeknisiRole(userid int) error {

	if err := us.userRepository.UpdateField(userid, "status", "unverified-teknisi"); err != nil {
		return err
	}

	return nil
}
