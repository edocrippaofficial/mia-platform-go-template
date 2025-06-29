package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetByIDRequest struct {
	ID     string `param:"id" validate:"required"`
	Name   string `query:"name"`
	ApiKey string `header:"x-api-key" validate:"required"`
}

type GetByIDResponse struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (ctrl *UserController) GetByIDHandler(c echo.Context) error {
	data := c.Get("data").(GetByIDRequest)

	user, err := ctrl.userService.GetByID(data.ID, data.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, GetByIDResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
