package main

import (
	"echotonic/middlewares"
	"echotonic/routes"
	"echotonic/spec"
	"log"
	"net/http"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = middlewares.NewValidator()
	openapi := &docs.OpenApi{
		OpenAPI: "3.0.1",
		Info: docs.InfoObject{
			Version: "1.0.0",
			Title:   "Echo Example API",
		},
	}

	routes.RegisterRoutes(e, openapi)
	spec.ExposeSwaggerUI(e, openapi)

	if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
