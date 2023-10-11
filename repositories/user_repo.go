package repositories

import (
	"praktikum_23/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) error
	Find() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user models.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) Find() ([]models.User, error) {
	var users []models.User

	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
