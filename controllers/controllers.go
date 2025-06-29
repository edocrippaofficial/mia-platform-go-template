package controllers

import (
	"echotonic/controllers/users"

	"echotonic/router"
)

type Controller interface {
	RegisterRoutes(r *router.Router)
}

func GetControllers() []Controller {
	return []Controller{
		users.NewUserController(),
	}
}
