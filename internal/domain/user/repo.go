package user

import (
	"gin/internal/database"
	"gin/internal/models"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where(&models.User{Email: email}).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByRefreshToken(refresh_token string) (*models.User, error) {
	var user models.User
	err := database.DB.Where(&models.User{RefreshToken: refresh_token}).First(&user).Error
	return &user, err
}

func (r *UserRepository) UpdateRefreshToken(userId uint, token string) error {
	return database.DB.Model(&models.User{}).Where("id = ?", userId).Update("refresh_token", token).Error
}
