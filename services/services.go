package services

import (
	userService "mia_template_service_name_placeholder/services/users"
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
