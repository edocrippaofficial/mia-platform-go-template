package router

import (
	"echotonic/middlewares"

	echoSwagger "github.com/TickLabVN/tonic/adapters/echo"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Handler *echo.Echo
	OpenAPI *docs.OpenApi
}

func NewRouter(logger *logrus.Logger) *Router {
	e := echo.New()
	e.HideBanner = true

	e.Validator = middlewares.NewValidator()
	e.Use(middlewares.RequestMiddlewareLogger(logger, []string{"/-/", "/documentation"}))

	openapi := &docs.OpenApi{
		OpenAPI: "3.0.1",
		Info: docs.InfoObject{
			Version: "1.0.0",
			Title:   "Echo Example API",
		},
	}

	router := &Router{
		Handler: e,
		OpenAPI: openapi,
	}

	addHealthRoutes(router)

	exposeSwaggerUI(e, openapi)

	return router
}

func RegisterRoute[Req any, Res any](r *Router, method string, path string, handler echo.HandlerFunc, opts ...docs.OperationObject) {
	route := r.Handler.Add(method, path, handler, middlewares.Bind[Req])
	echoSwagger.AddRoute[Req, Res](r.OpenAPI, route, opts...)
}
