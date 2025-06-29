package users

import (
	"errors"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService interface {
	GetByID(id string, name string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetByID(id string, name string) (*User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}

	return user, nil
}
