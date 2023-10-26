package repository

import (
	"errors"
	"terhandle/internal/app/model"
	"terhandle/internal/features/request-teknisi/entity"

	"gorm.io/gorm"
)

type userRequestRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.UserRequestRepositoryInterface {
	return &userRequestRepository{
		db: db,
	}
}

func (ur *userRequestRepository) Insert(data entity.Core) error {

	user := entity.UserCoreToUserModel(data)
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRequestRepository) SelectByIdAndRole(userid, teknisiid int, role_user, role_teknisi string) error {
	var userUser model.Users
	var userTeknisi model.Users
	roleUser := ur.db.Where("id = ? AND role = ?", userid, role_user).First(&userUser)
	if roleUser.Error != nil {
		return roleUser.Error
	}

	roleTeknisi := ur.db.Where("id = ? AND role = ?", teknisiid, role_teknisi).First(&userTeknisi)
	if roleTeknisi.Error != nil {
		return roleTeknisi.Error
	}

	return nil
}

func (ur *userRequestRepository) SelectById(userid int) ([]entity.Core, error) {
	var userHistory []model.RequestTeknisi

	if err := ur.db.Preload("Foto").Find(&userHistory).Error; err != nil {
		return nil, errors.New("gagal mengambil data")
	}

	var usersHistory []entity.Core = entity.UserModelToUserCoreList(userHistory)
	return usersHistory, nil
}

func (ur *userRequestRepository) Update(userid int, data entity.Core) error {
	user := entity.UserCoreToUserModel(data)

	if err := ur.db.Where("id = ?", userid).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRequestRepository) UpdateField(userid int, field string, value string) error {
	var user model.Users
	if err := ur.db.Where("id = ? AND (NIK != 0 AND No_telp != 0 AND Alamat != '')", userid).First(&user).Error; err != nil {
		return errors.New("lengkapi dahulu NIK, No_telp, Alamat")
	}

	if err := ur.db.Model(&model.Users{}).Where("id = ?", userid).Update(field, value).Error; err != nil {
		return err
	}

	return nil
}
