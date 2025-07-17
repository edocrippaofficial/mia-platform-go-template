package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Bind[Req any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data Req
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input", "message": err.Error()})
		}
		err := (&echo.DefaultBinder{}).BindHeaders(c, &data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid headers", "message": err.Error()})
		}
		if err := c.Validate(&data); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input", "validation": validateErrorMapping(err), "message": err.Error()})
		}
		c.Set("data", data)
		return next(c)
	}
}
