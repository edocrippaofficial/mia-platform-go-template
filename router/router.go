package router

import (
	"echotonic/middlewares"

	echoSwagger "github.com/TickLabVN/tonic/adapters/echo"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo    *echo.Echo
	OpenAPI *docs.OpenApi
}

func NewRouter(e *echo.Echo, openapi *docs.OpenApi) *Router {
	return &Router{
		Echo:    e,
		OpenAPI: openapi,
	}
}

func RegisterRoute[Req any, Res any](r *Router, method string, path string, handler echo.HandlerFunc, opts ...docs.OperationObject) {
	route := r.Echo.Add(method, path, handler, middlewares.Bind[Req])
	echoSwagger.AddRoute[Req, Res](r.OpenAPI, route, opts...)
}
