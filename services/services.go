package services

import (
	userService "echotonic/services/users"
)

// Services holds all the application dependencies
type Services struct {
	UserService userService.UserService
}

func NewServices() *Services {
	userRepo := userService.NewInMemoryUserRepository()
	userSvc := userService.NewUserService(userRepo)

	return &Services{
		UserService: userSvc,
	}
}
