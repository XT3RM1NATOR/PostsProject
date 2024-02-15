package service

import (
	"github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/XT4RM1NATOR/PostsProject/user_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/user_service/util"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(req *user_service.CreateUserRequest) (user_service.CreateUserResponse, error) {
	hashedPassword, err := util.GeneratePasswordHash(req.Password)
	if err != nil {
		return user_service.CreateUserResponse{}, err
	}

	err = s.Repo.CreateUser(req.Username, hashedPassword, req.Role.String())
	if err != nil {
		return user_service.CreateUserResponse{}, err
	}

	return user_service.CreateUserResponse{}, err
}
