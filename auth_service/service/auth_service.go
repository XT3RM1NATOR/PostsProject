package service

import (
	"context"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/util"
	userPb "github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/XT4RM1NATOR/PostsProject/shared"
	"log"
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

	userClient, err := shared.GetUserServiceClient()
	if err != nil {
		log.Println("Err creating a connection with user client")
	}
	req := &userPb.CheckUserByPropertyRequest{
		Username: username,
		Password: password,
	}

	user, err := userClient.CheckUserByProperty(ctx, req)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := util.GenerateTokens(int(user.Id), string(user.Role))
	if err != nil {
		return "", "", err
	}

	daysLeft, err := strconv.Atoi(os.Getenv("DAYS_REFRESH_TOKEN"))
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(daysLeft))
	err = s.Repo.CreateSession(int(user.Id), refreshToken, expiresAt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Logout(_ context.Context, refreshToken string) error {
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

func (s *AuthService) Register(ctx context.Context, username, password string, role string) (string, string, error) {

	req := &userPb.CreateUserRequest{
		Username: username,
		Password: password,
		Role:     role,
	}

	userClient, err := shared.GetUserServiceClient()
	if err != nil {
		log.Println("Err creating a connection with user client")
	}

	user, err := userClient.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("i am user client err")
		return "", "", err
	}

	accessToken, refreshToken, err := util.GenerateTokens(int(user.Id), role)
	if err != nil {
		return "", "", err
	}

	daysLeft, err := strconv.Atoi(os.Getenv("DAYS_REFRESH_TOKEN"))
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(daysLeft))
	err = s.Repo.CreateSession(int(user.Id), refreshToken, expiresAt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
