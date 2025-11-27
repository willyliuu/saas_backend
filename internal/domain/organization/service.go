package organization

import "gin/internal/models"

type OrganizationService struct {
	Repo *OrganizationRepository
}

func (s *OrganizationService) Create(dto CreateOrganizationDTO, ownerID uint) (*models.Organization, error) {
	organization := &models.Organization{
		Name:    dto.Name,
		OwnerID: ownerID,
	}

	err := s.Repo.Create(organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) FindAll() (*[]models.Organization, error) {
	return s.Repo.FindAll()
}

func (s *OrganizationService) FindByID(orgID uint) (*models.Organization, error) {
	return s.Repo.FindByID(orgID)
}

func (s *OrganizationService) Update(id uint, dto UpdateOrganizationDTO) (*models.Organization, error) {
	// check if the organization exists
	organization, err := s.Repo.FindByID(id)
	if err != nil {
		return organization, err
	}

	// update field
	organization.Name = dto.Name

	if err := s.Repo.Update(organization); err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
