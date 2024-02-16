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

	id, err := s.Repo.CreateUser(req.Username, hashedPassword, req.Role)
	if err != nil {
		return user_service.CreateUserResponse{}, err
	}

	return user_service.CreateUserResponse{
		Id:   id,
		Role: req.Role,
	}, err
}

func (s *UserService) GetUserById(req *user_service.GetUserByIdRequest) (user_service.UserResponse, error) {
	user, err := s.Repo.GetUserByID(int(req.Id))
	if err != nil {
		return user_service.UserResponse{}, err
	}

	return *user, nil
}

func (s *UserService) GetUsers() (user_service.UsersResponse, error) {
	users, err := s.Repo.GetUsers()
	return users, err
}

func (s *UserService) UpdateUser(req *user_service.UpdateUserRequest) error {
	err := s.Repo.UpdateUser(int(req.Id), &req.Username, &req.Role)
	return err
}

func (s *UserService) DeleteUser(req *user_service.DeleteUserRequest) error {
	err := s.Repo.DeleteUser(int(req.Id))
	return err
}

func (s *UserService) CheckUserByProperty(req *user_service.CheckUserByPropertyRequest) (user_service.CheckUserByPropertyResponse, error) {
	user, err := s.Repo.GetUserByProperty(req.Username)
	if err != nil {
		return user_service.CheckUserByPropertyResponse{}, err
	}
	err = util.ComparePasswordAndHash(req.Password, user.PasswordHash)
	if err != nil {
		return user_service.CheckUserByPropertyResponse{}, err
	}
	return user_service.CheckUserByPropertyResponse{
		Id:   user.Id,
		Role: user.Role,
	}, nil
}
