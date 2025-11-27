package invite

import (
	"errors"
	"time"

	"gin/internal/models"

	"github.com/google/uuid"
)

// func invitemember

type InviteService struct {
	Repo *InviteRepository
}

func (s *InviteService) InviteMember(orgID uint, dto CreateInviteDTO) (*models.Invite, error) {
	existingInvite, _ := s.Repo.FindByEmail(dto.Email)
	if existingInvite.ID != 0 {
		return nil, errors.New("user already invited")
	}

	token := uuid.NewString()

	invite := &models.Invite{
		OrganizationID: orgID,
		Email:          dto.Email,
		Role:           dto.Role,
		Status:         "pending",
		Token:          token,
		ExpiresAt:      time.Now().Add(48 * time.Hour),
	}
	if err := s.Repo.Create(invite); err != nil {
		return nil, err
	}

	return invite, nil
}

func (s *InviteService) GetInviteByToken(token string) (*models.Invite, error) {
	return s.Repo.FindByToken(token)
}

func (s *InviteService) AcceptInvite(userID uint, dto *models.Invite) (*models.UserOrganization, error) {
	// check expired invite
	if time.Now().After(dto.ExpiresAt) {
		return nil, errors.New("invite expired")
	}

	// create membership
	userOrg := &models.UserOrganization{
		UserID:         userID,
		OrganizationID: dto.OrganizationID,
		Role:           dto.Role,
	}

	if err := s.Repo.CreateUserOrganization(userOrg); err != nil {
		return nil, err
	}

	// update invite to be accepted
	dto.Status = "accepted"
	if err := s.Repo.UpdateInvite(dto); err != nil {
		return nil, err
	}

	return userOrg, nil
}
