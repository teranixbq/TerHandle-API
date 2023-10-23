package repository

import (
	"errors"
	"terhandle/internal/app/model"
	"terhandle/internal/features/user/entity"
	"terhandle/internal/utils/helper"
	"terhandle/internal/utils/jwt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Insert(data entity.Core) error {

	securePassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return err
	}

	user := entity.UserCoreToUserModel(data)
	user.Password = securePassword

	if err = ur.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Login(email, password string) (entity.Core, string, error) {
	var user model.Users

	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entity.Core{}, "", result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Core{}, "", errors.New("pengguna tidak ditemukan")
	}

	if compare := helper.CompareHash(user.Password, password); !compare {
		return entity.Core{}, "", errors.New("kata sandi salah")
	}


	token, errToken := jwt.CreateToken(int(user.ID), user.Role)
	if errToken != nil {
		return entity.Core{}, "", errToken
	}

	userCoreData := entity.UserModelToUserCore(user)
	return userCoreData, token, nil
}

func (ur *userRepository) Update(userid int, data entity.Core) error {
	user := entity.UserCoreToUserModel(data)

	if err := ur.db.Where("id = ?", userid).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UpdateField(userid int, field string, value string) error {
    var user model.Users
    if err := ur.db.Where("id = ? AND (NIK != 0 AND No_telp != 0 AND Address != '')", userid).First(&user).Error; err != nil {
        return errors.New("lengkapi dahulu NIK, No_telp, Alamat")
    }

    if err := ur.db.Model(&model.Users{}).Where("id = ?", userid).Update(field, value).Error; err != nil {
        return err
    }

    return nil
}

