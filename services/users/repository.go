package users

import (
	"fmt"
)

type UserRepository interface {
	FindByID(id string) (*User, error)
}

type inMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() UserRepository {
	users := map[string]*User{
		"1": {
			ID:    "1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		"2": {
			ID:    "2",
			Name:  "Jane Smith",
			Email: "jane.smith@example.com",
		},
		"3": {
			ID:    "3",
			Name:  "Bob Johnson",
			Email: "bob.johnson@example.com",
		},
	}

	return &inMemoryUserRepository{
		users: users,
	}
}

func (r *inMemoryUserRepository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
