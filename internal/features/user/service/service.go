package service

import (
	"errors"
	"fmt"
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
	fmt.Println(payload.Email,payload.Password,payload.Nama)
	if payload.Email == "" || payload.Password == "" || payload.Nama == "" {
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

func (us *userService) CreateUpdateDetail(id_user uint, data entity.Core) error {

	if err := us.userRepository.Update(id_user, data); err != nil {
		return err
	}

	return nil
}

func (us *userService) RequestTeknisiRole(id_user uint) error {

	if err := us.userRepository.UpdateField(id_user, "role", "teknisi"); err != nil {
		return err
	}

	return nil
}

func (us *userService) GetAll() ([]entity.Core, error) {
	users, err := us.userRepository.SelectAll()
	if err != nil {
		return nil, err
	}
	return users, err
}

func (us *userService) GetById(id_user uint) ([]entity.Core, error) {
	users, err := us.userRepository.SelectById(id_user)
	if err != nil {
		return nil, err
	}
	return users, err
}


func (us *userService) DeleteById(id_user uint) error {
	err := us.userRepository.DeleteById(id_user)
	if err != nil {
		return nil
	}
	return nil
}
