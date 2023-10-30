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

func (ur *userRepository) Update(id_user uint, data entity.Core) error {
	user := entity.UserCoreToUserModel(data)

	if err := ur.db.Where("id = ?", id_user).Updates(&user).Error; err != nil {
		return errors.New("profesi tidak ada")
	}

	return nil
}

func (ur *userRepository) UpdateField(id_user uint, field string, value string) error {
	var user model.Users
	if err := ur.db.Where("id = ? AND (NIK != 0 AND No_telp != 0 AND Alamat != '')", id_user).First(&user).Error; err != nil {
		return errors.New("lengkapi dahulu NIK, No_telp, Alamat")
	}

	if err := ur.db.Model(&model.Users{}).Where("id = ?", id_user).Update(field, value).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) SelectAll() ([]entity.Core, error) {
	var usersData []model.Users
	err := ur.db.Where("role = ?", "teknisi").Find(&usersData)
	if err.Error != nil {
		return nil, err.Error
	}
	var usersCoreAll []entity.Core = entity.UserModelToUserCoreList(usersData)
	return usersCoreAll, nil
}

func (ur *userRepository) SelectUserById(id_user uint, role string) (entity.Core, error) {
	var usersData model.Users
	err := ur.db.Where("id = ? AND role = ?", id_user, role).Find(&usersData)
	if err.Error != nil {
		return entity.Core{}, err.Error
	}
	var usersCoreAll entity.Core = entity.UserModelToUserCore(usersData)

	return usersCoreAll, nil
}

func (ur *userRepository) SelectByIdWithFeedback(id_user uint) ([]entity.Core, error) {
	var usersData []model.Users
	err := ur.db.Preload("Feedback").Where("id = ? AND role = ?", id_user, "teknisi").Find(&usersData)
	if err.Error != nil {
		return nil, err.Error
	}
	var usersCoreAll []entity.Core = entity.UserModelToUserCoreList(usersData)

	return usersCoreAll, nil
}

func (ur *userRepository) DeleteById(id_user uint) error {
	var user model.Users
	if err := ur.db.First(&user, id_user).Error; err != nil {
		return err
	}

	if err := ur.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
