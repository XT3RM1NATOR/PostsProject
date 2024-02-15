package main

import (
	"github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/XT4RM1NATOR/PostsProject/user_service/service"
)

type server struct {
	userService *service.UserService
	user_service.UnimplementedUserServiceServer
}

func (s *server) CreateUser(req *user_service.CreateUserRequest) (*user_service.CreateUserResponse, error) {
	res, err := s.userService.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
