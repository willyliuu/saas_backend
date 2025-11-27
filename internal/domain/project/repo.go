package project

import (
	"gin/internal/database"
	"gin/internal/models"
)

type ProjectRepository struct{}

func (r *ProjectRepository) Create(project *models.Project) error {
	return database.DB.Create(project).Error
}

func (r *ProjectRepository) FindByOrgID(orgID uint) ([]*models.Project, error) {
	var projects []*models.Project
	err := database.DB.Where(&models.Project{OrganizationID: orgID}).Find(&projects).Error

	return projects, err
}

func (r *ProjectRepository) FindByID(id uint) (*models.Project, error) {
	var project *models.Project
	err := database.DB.Where(&models.Project{ID: id}).First(&project).Error

	return project, err
}

func (r *ProjectRepository) FindByOrgIDAndName(orgID uint, name string) (*models.Project, error) {
	var project *models.Project
	err := database.DB.Where(&models.Project{OrganizationID: orgID, Name: name}).First(&project).Error

	return project, err
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return database.DB.Save(project).Error
}

func (r *ProjectRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Project{}, id).Error
}
