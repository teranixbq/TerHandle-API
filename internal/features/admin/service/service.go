package service

import (
	"terhandle/internal/features/admin/entity"
)

type AdminService struct {
	AdminRepository entity.AdminRepositoryInterface
}

func NewUserService(Ad entity.AdminRepositoryInterface) entity.AdminServiceInterface {
	return &AdminService{
		AdminRepository: Ad,
	}
}

func (Ad *AdminService) Create(payload entity.CoreProfesi) error {

	err := Ad.AdminRepository.Insert(payload)
	if err != nil {
		return err
	}

	return nil
}

func (Ad *AdminService) GetAll() ([]entity.CoreProfesi, error) {
	profesiAll, err := Ad.AdminRepository.SelectAll()
	if err != nil {
		return nil, err
	}

	return profesiAll, err
}

func (Ad *AdminService) Update(id_profesi uint, payload entity.CoreProfesi) error {
	err := Ad.AdminRepository.Update(id_profesi, payload)
	if err != nil {
		return err
	}

	return nil
}

func (Ad *AdminService) Delete(id_profesi uint) error {
	err := Ad.AdminRepository.Delete(id_profesi)
	if err != nil {
		return err
	}

	return nil
}

func (Ad *AdminService) InsertBiaya(payload entity.CoreTransport) error {
	err := Ad.AdminRepository.InsertBiaya(payload)
	if err != nil {
		return err
	}

	return err
}

func (Ad *AdminService) UpdateBiaya(id uint, payload entity.CoreTransport) error {
	err := Ad.AdminRepository.UpdateBiaya(id, payload)
	if err != nil {
		return err
	}

	return err
}
