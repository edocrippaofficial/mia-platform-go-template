package users

import (
	"echotonic/router"

	"github.com/TickLabVN/tonic/core/docs"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

var tags = []string{"users"}

func (c *UserController) RegisterRoutes(r *router.Router) {
	router.RegisterRoute[GetByIDRequest, GetByIDResponse](r,
		"GET", "/users/:id",
		GetByIDHandler,
		docs.OperationObject{OperationId: "get-user-by-id", Tags: tags, Summary: "Get user by ID"},
	)
}
