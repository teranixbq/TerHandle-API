package service

import (
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

func (us *userService) ClaimRequest(id_request int, data entity.Core) error {

	if err := us.userRepository.UpdateClaims(id_request, data); err != nil {
		return err
	}

	if err := us.userRepository.UpdateField(id_request, "status", "konfirmasi biaya"); err != nil {
		return err
	}

	return nil
}

func (us *userService) KonfirmasiBiaya(id_user, id_request int, data entity.Core) error {

	if err := us.userRepository.UpdateStatusKonfirmBiaya(id_user, id_request, data); err != nil {
		return err
	}

	if err := us.userRepository.UpdateField(id_request, "status", "diproses"); err != nil {
		return err
	}

	return nil
}

func (us *userService) BatalkanRequest(id_user, id_request int, data entity.Core) error {

	if err := us.userRepository.UpdateStatusBatal(id_user, id_request, data); err != nil {
		return err
	}
	if err := us.userRepository.UpdateField(id_request, "status", "dibatalkan"); err != nil {
		return err
	}

	return nil
}

func (us *userService) SelesaikanRequest(id_user, id_request int, data entity.Core) error {

	if err := us.userRepository.UpdateStatusSelesai(id_user, id_request, data); err != nil {
		return err
	}
	if err := us.userRepository.UpdateField(id_request, "status", "selesai"); err != nil {
		return err
	}

	return nil
}
