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

func (ur *userRequestRepository) SelectAllById(userid int) ([]entity.Core, error) {
	var userHistory []model.RequestTeknisi

	if err := ur.db.Preload("Foto").Where("users_id = ? OR teknisi_id = ?", userid,userid).Find(&userHistory).Error; err != nil {
		return nil, errors.New("gagal mengambil data")
	}

	var usersHistory []entity.Core = entity.UserModelToUserCoreList(userHistory)
	return usersHistory, nil
}

func (ur *userRequestRepository) SelectById(user_id, id int) ([]entity.Core, error) {
	var userHistory []model.RequestTeknisi

	if err := ur.db.Preload("Foto").Where("(users_id = ? OR teknisi_id = ?) AND id = ?",user_id, user_id, id).Find(&userHistory).Error; err != nil {
		return nil, errors.New("gagal mengambil data")
	}

	var usersHistory = entity.UserModelToUserCoreList(userHistory)
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

	if err := ur.db.Model(&model.RequestTeknisi{}).Where("id = ?", userid).Update(field, value).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRequestRepository) UpdateClaims(id int, data entity.Core) error {
	request := entity.UserCoreToUserModel(data)
	if err := ur.db.First(&request, id).Error; err != nil {
		return err
	}

	request.Biaya = data.Biaya
	request.Diproses = data.Diproses
	request.Dibatalkan = data.Dibatalkan
	request.Menunggu_konfirmasi = data.Menunggu_konfirmasi

	if err := ur.db.Save(&request).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRequestRepository) UpdateStatusClaims(id int, data entity.Core) error {
	request := entity.UserCoreToUserModel(data)
	if err := ur.db.First(&request, id).Error; err != nil {
		return err
	}

	request.Diproses = data.Diproses
	request.Konfirmasi_biaya = data.Konfirmasi_biaya
	request.Dibatalkan = data.Dibatalkan
	request.Selesai = data.Selesai

	if err := ur.db.Save(&request).Error; err != nil {
		return err
	}

	return nil
}
