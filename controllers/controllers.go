package controllers

import (
	"echotonic/controllers/users"
	"echotonic/services"

	"echotonic/router"
)

type Controller interface {
	RegisterRoutes(r *router.Router)
}

func NewControllers(svcs *services.Services) []Controller {
	return []Controller{
		users.NewUserController(svcs),
	}
}
