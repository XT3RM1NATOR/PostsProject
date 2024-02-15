package service

import (
	"context"
	"errors"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/util"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	Repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	err = util.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := util.GenerateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	daysLeft, err := strconv.Atoi(os.Getenv("DAYS_REFRESH_TOKEN"))
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(daysLeft))
	err = s.Repo.CreateSession(user.ID, refreshToken, expiresAt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	sessionID, err := s.Repo.GetSessionByRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Register(ctx context.Context, username, email, password string, role string) error {
	hashedPassword, err := util.GeneratePasswordHash(password)
	if err != nil {
		return err
	}

	err = s.Repo.CreateUser(username, email, hashedPassword, role)
	if err != nil {
		return err
	}

	return nil
}
