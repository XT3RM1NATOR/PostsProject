package main

import (
	"context"
	"log"
	"net"

	"github.com/XT4RM1NATOR/PostsProject/auth_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/service"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/util"
	"github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
	"google.golang.org/grpc"
)

type server struct {
	authService *service.AuthService
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
	accessToken, err := s.authService.RefreshAccessToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &auth_service.AccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db := // Initialize your database connection
	repo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(repo)
	s := grpc.NewServer()
	auth_service.RegisterAuthServiceServer(s, &server{authService: authService})
	log.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
