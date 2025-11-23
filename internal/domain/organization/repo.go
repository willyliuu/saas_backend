package organization

import (
	"gin/internal/database"
	"gin/internal/models"
)

type OrganizationRepository struct{}

func (r *OrganizationRepository) Create(organization *models.Organization) error {
	return database.DB.Create(organization).Error
}

func (r *OrganizationRepository) FindAll() (*[]models.Organization, error) {
	var organizations []models.Organization
	err := database.DB.Find(&organizations).Error

	return &organizations, err
}

func (r *OrganizationRepository) FindByID(orgID uint) (*models.Organization, error) {
	var organization models.Organization
	err := database.DB.Where(&models.Organization{ID: orgID}).First(&organization).Error

	return &organization, err
}
