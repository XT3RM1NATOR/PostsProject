package main

import (
	"context"
	"github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/XT4RM1NATOR/PostsProject/user_service/service"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	userService *service.UserService
	user_service.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *user_service.CreateUserRequest) (*user_service.CreateUserResponse, error) {
	res, err := s.userService.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Server) GetUserById(ctx context.Context, req *user_service.GetUserByIdRequest) (*user_service.UserResponse, error) {
	user, err := s.userService.GetUserById(req)
	return &user, err
}

func (s *Server) GetUsers(_ context.Context, _ *empty.Empty) (*user_service.UsersResponse, error) {
	users, err := s.userService.GetUsers()
	return &users, err
}

func (s *Server) UpdateUser(_ context.Context, req *user_service.UpdateUserRequest) (*empty.Empty, error) {
	err := s.userService.UpdateUser(req)
	return nil, err
}

func (s *Server) DeleteUser(_ context.Context, req *user_service.DeleteUserRequest) (*empty.Empty, error) {
	err := s.userService.DeleteUser(req)
	return nil, err
}

func (s *Server) CheckUserByProperty(_ context.Context, req *user_service.CheckUserByPropertyRequest) (*user_service.CheckUserByPropertyResponse, error) {
	user, err := s.userService.CheckUserByProperty(req)
	return &user, err
}
