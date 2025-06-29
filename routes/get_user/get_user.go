package getuser

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	ID     string `param:"id" validate:"required"`
	Name   string `query:"name"`
	ApiKey string `header:"x-api-key" validate:"required"`
}

type Response struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func Handler(c echo.Context) error {
	data := c.Get("data").(Request)
	return c.JSON(http.StatusOK, Response{
		ID:    data.ID,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})
}
