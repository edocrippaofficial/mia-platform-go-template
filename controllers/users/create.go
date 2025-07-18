package users

import (
	"mia_template_service_name_placeholder/services/users"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	ID       string `query:"id"`
	Name     string `json:"name" validate:"required"`
	Metadata *struct {
		Age   int    `json:"age"`
		Email string `json:"email" validate:"required,email"`
	} `json:"metadata,omitempty" validate:"omitempty"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

func (ctrl *UserController) CreateHandler(c echo.Context) error {
	data := c.Get("data").(CreateRequest)
	fmt.Println("CreateHandler called with data:", data)

	user := &users.User{
		Name: data.Name,
	}

	if data.Metadata != nil {
		user.Email = data.Metadata.Email
	}

	response, err := ctrl.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, CreateResponse{
		ID: response.ID,
	})
}
