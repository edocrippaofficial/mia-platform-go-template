package main

import (
	"echotonic/middlewares"
	"echotonic/routes"
	"echotonic/spec"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = middlewares.NewValidator()

	openapi := &docs.OpenApi{
		OpenAPI: "3.0.1",
		Info: docs.InfoObject{
			Version: "1.0.0",
			Title:   "Echo Example API",
		},
	}

	router := routes.NewRouter(e, openapi)
	router.RegisterRoutes()

	spec.ExposeOpenAPI(e, openapi)

	e.Logger.Fatal(e.Start(":3000"))
}
