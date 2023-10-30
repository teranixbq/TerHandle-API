package repository

import (
	"errors"
	"terhandle/internal/app/model"
	"terhandle/internal/features/admin/entity"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.AdminRepositoryInterface {
	return &AdminRepository{
		db: db,
	}
}

func (ur *AdminRepository) Insert(data entity.CoreProfesi) error {
    dataProfesi := model.Profesi {}

    if err := ur.db.Where("nama_profesi = ?", data.NamaProfesi).First(&dataProfesi).Error; err == nil {
        return errors.New("profesi sudah ada")

    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }

	profesi := entity.CoreProfesiToModelProfesi(data)
	if err := ur.db.Create(&profesi).Error; err != nil {
		return err
	}

	return nil
}

func (ur *AdminRepository) SelectAll() ([]entity.CoreProfesi, error) {
	profesi := []model.Profesi{}

	err := ur.db.Find(&profesi)
	if err.Error != nil {
		return nil, err.Error
	}

	var profesiAll []entity.CoreProfesi = entity.ProfesiModelToCoreProfesiList(profesi)
	return profesiAll, nil
}

func (ur *AdminRepository) Update(id_profesi uint, data entity.CoreProfesi) error {
    profesi := model.Profesi{}

    if err := ur.db.First(&profesi, id_profesi).Error; err != nil {
        return err
    }

    profesi.NamaProfesi = data.NamaProfesi
    if err := ur.db.Save(&profesi).Error; err != nil {
        return err
    }

    return nil
}


func (ur *AdminRepository) Delete(id_profesi uint) error {
	profesi := model.Profesi {}

	result := ur.db.Where("id = ?", id_profesi).Delete(&profesi)
	if result.Error != nil {
		return errors.New("profesi tidak ditemukan")
	}

	if result.RowsAffected == 0 {
		return errors.New("profesi tidak ditemukan")
	}

	return nil
}


func (ur *AdminRepository) InsertBiaya(data entity.CoreTransport) error {
    var count int64
    if err := ur.db.Model(&model.Transport{}).Where("id = ?", 1).Count(&count).Error; err != nil {
        return err
    }

    if count > 1 {
        return errors.New("data melebihi")
    } else if count == 1 {
        return errors.New("data sudah ada")
    }

    biaya := entity.CoreTransportToModelTransport(data)
    if err := ur.db.Create(&biaya).Error; err != nil {
        return err
    }

    return nil
}
