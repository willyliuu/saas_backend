package project

import (
	"errors"

	"gin/internal/models"
)

type ProjectService struct {
	Repo *ProjectRepository
}

func (s *ProjectService) Create(orgID uint, dto CreateProjectDTO) (*models.Project, error) {
	existingProject, _ := s.Repo.FindByOrgIDAndName(orgID, dto.Name)
	if existingProject.ID != 0 {
		return nil, errors.New("project already exists")
	}

	project := &models.Project{
		Name:           dto.Name,
		Description:    dto.Description,
		Archieved:      dto.Archieved,
		OrganizationID: orgID,
	}

	if err := s.Repo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) FindByOrgID(orgID uint) ([]*models.Project, error) {
	return s.Repo.FindByOrgID(orgID)
}

func (s *ProjectService) FindByID(id uint) (*models.Project, error) {
	return s.Repo.FindByID(id)
}

func (s *ProjectService) Update(id uint, dto UpdateProjectDTO) (*models.Project, error) {
	project, err := s.Repo.FindByID(id)
	if err != nil {
		return project, err
	}

	project.Name = dto.Name
	project.Description = dto.Description
	project.Archieved = dto.Archieved

	if err := s.Repo.Update(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
