package controllers

import (
	"echotonic/controllers/users"

	"echotonic/routes"
)

type Controller interface {
	RegisterRoutes(r *routes.Router)
}

func GetControllers() []Controller {
	return []Controller{
		users.NewUserController(),
	}
}
