package user

import (
	"errors"

	"gin/internal/models"
	"gin/internal/token"
	"gin/internal/utils"
)

type UserService struct {
	Repo  *UserRepository
	Token *token.TokenMaker
}

func (s *UserService) Register(dto RegisterDTO) error {
	hashedPassword, _ := utils.HashPassword(dto.Password)

	user := &models.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: hashedPassword,
	}

	return s.Repo.Create(user)
}

func (s *UserService) Login(dto LoginDTO) (string, string, error) {
	user, err := s.Repo.FindByEmail(dto.Email)
	if err != nil || !utils.CheckPassword(dto.Password, user.PasswordHash) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, _ := s.Token.CreateAccessToken(user.ID)
	refreshToken, _ := s.Token.CreateRefreshToken(user.ID)

	s.Repo.UpdateRefreshToken(user.ID, refreshToken)

	return accessToken, refreshToken, nil
}

func (s *UserService) Refresh(old string) (string, string, error) {
	user, err := s.Repo.FindByRefreshToken(old)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	accessToken, _ := s.Token.CreateAccessToken(user.ID)
	refreshToken, _ := s.Token.CreateRefreshToken(user.ID)

	s.Repo.UpdateRefreshToken(user.ID, refreshToken)

	return accessToken, refreshToken, nil
}

func (s *UserService) GetMe(userID uint) (*models.User, error) {
	return s.Repo.FindByID(userID)
}
