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

func GetByIDHandler(c echo.Context) error {
	data := c.Get("data").(GetByIDRequest)

	name := "John Doe"
	if data.Name != "" {
		name = data.Name
	}

	return c.JSON(http.StatusOK, GetByIDResponse{
		ID:    data.ID,
		Name:  name,
		Email: "john.doe@example.com",
	})
}
