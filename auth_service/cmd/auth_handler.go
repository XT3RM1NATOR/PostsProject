package main

import (
	"context"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/service"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/util"
	"github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
)

type server struct {
	authService *service.AuthService
	auth_service.UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, req *auth_service.AuthRequest) (*auth_service.AuthResponse, error) {
	accessToken, refreshToken, err := s.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &auth_service.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *server) Register(ctx context.Context, req *auth_service.User) (*auth_service.RegistrationResponse, error) {
	err := s.authService.Register(ctx, req.Username, req.Email, req.Password, req.Role)
	if err != nil {
		return nil, err
	}
	return &auth_service.RegistrationResponse{
		Response: "User registered successfully",
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *auth_service.UserRequest) (*auth_service.UserResponse, error) {
	accessToken := req.AccessToken

	claims, err := util.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.authService.Repo.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	return &auth_service.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func (s *server) GetAccessToken(ctx context.Context, req *auth_service.AccessTokenRequest) (*auth_service.AccessTokenResponse, error) {
	sessionID, err := s.authService.Repo.GetSessionByRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.authService.Repo.GetUserByID(sessionID)
	if err != nil {
		return nil, err
	}

	accessToken, err := util.GenerateToken(user.ID, user.Role, util.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	return &auth_service.AccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
