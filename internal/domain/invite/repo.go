package invite

import (
	"gin/internal/database"
	"gin/internal/models"
)

// findBy Email, create invite, delete by token

type InviteRepository struct{}

func (r *InviteRepository) Create(invite *models.Invite) error {
	return database.DB.Create(invite).Error
}

func (r *InviteRepository) FindByEmail(email string) (*models.Invite, error) {
	var invite *models.Invite
	err := database.DB.Where(&models.Invite{Email: email}).First(&invite).Error

	return invite, err
}

func (r *InviteRepository) FindByToken(token string) (*models.Invite, error) {
	var invite *models.Invite
	err := database.DB.Where(&models.Invite{Token: token}).First(&invite).Error

	return invite, err
}

func (r *InviteRepository) UpdateInvite(invite *models.Invite) error {
	return database.DB.Save(invite).Error
}

func (r *InviteRepository) DeleteByToken(token string) error {
	return database.DB.Where(&models.Invite{Token: token}).Delete(&models.Invite{}).Error
}

func (r *InviteRepository) CreateUserOrganization(userOrganization *models.UserOrganization) error {
	return database.DB.Create(userOrganization).Error
}
