package router

import (
	"net/http"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

type HealthRequest struct{}

type HealthResponse struct {
	Status string `json:"status"`
}

var tags = []string{"health"}

// registers the health check routes to the router.
func addHealthRoutes(r *Router) {
	RegisterRoute[HealthRequest, HealthResponse](r,
		"GET", "/-/ready",
		readyHandler,
		docs.OperationObject{OperationId: "ready", Tags: tags, Summary: "Readiness check"},
	)

	RegisterRoute[HealthRequest, HealthResponse](r,
		"GET", "/-/healthz",
		healthzHandler,
		docs.OperationObject{OperationId: "healthz", Tags: tags, Summary: "Health check"},
	)

}

func readyHandler(c echo.Context) error {
	response := HealthResponse{Status: "OK"}
	return c.JSON(http.StatusOK, response)
}

func healthzHandler(c echo.Context) error {
	response := HealthResponse{Status: "OK"}
	return c.JSON(http.StatusOK, response)
}
