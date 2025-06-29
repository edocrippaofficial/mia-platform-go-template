package users

import (
	"echotonic/router"
	"echotonic/services"
	userService "echotonic/services/users"

	"github.com/TickLabVN/tonic/core/docs"
)

type UserController struct {
	userService userService.UserService
}

func NewUserController(svc *services.Services) *UserController {
	return &UserController{
		userService: svc.UserService,
	}
}

var tags = []string{"users"}

func (c *UserController) RegisterRoutes(r *router.Router) {
	router.RegisterRoute[GetByIDRequest, GetByIDResponse](r,
		"GET", "/users/:id",
		c.GetByIDHandler,
		docs.OperationObject{OperationId: "get-user-by-id", Tags: tags, Summary: "Get user by ID"},
	)
}
