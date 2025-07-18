package controllers

import (
	"mia_template_service_name_placeholder/controllers/users"
	"mia_template_service_name_placeholder/router"
	"mia_template_service_name_placeholder/services"
)

type Controller interface {
	RegisterRoutes(r *router.Router)
}

func NewControllers(svcs *services.Services) []Controller {
	return []Controller{
		users.NewUserController(svcs),
	}
}
