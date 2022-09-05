package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type userService struct {
	repository repository.UserRepository
}

type UserService interface {
	ListUsers() (*[]domain.User, error)
	CreateUser(*domain.User) (*domain.User, error)
	ReadUser(ID string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(ID string) error
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) ListUsers() (*[]domain.User, error) {
	return s.repository.ListUsers()
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	return s.repository.CreateUser(user)
}

func (s *userService) ReadUser(ID string) (*domain.User, error) {
	return s.repository.ReadUser(ID)
}

func (s *userService) UpdateUser(user *domain.User) (*domain.User, error) {
	return s.repository.UpdateUser(user)
}

func (s *userService) DeleteUser(ID string) error {
	return s.repository.DeleteUser(ID)
}
