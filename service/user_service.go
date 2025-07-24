package service

import (
	"user-service/model"
	"user-service/repository"
)

type UserService interface {
	GetUser(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id int64) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
