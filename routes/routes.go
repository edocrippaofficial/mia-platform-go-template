package routes

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

type HandlerWithTypes[Req any, Res any] struct {
	Handler echo.HandlerFunc
	Method  string
	Path    string
	Options []docs.OperationObject
}

func RegisterRoute[Req any, Res any](r *Router, method string, path string, handler echo.HandlerFunc, opts ...docs.OperationObject) {
	h := HandlerWithTypes[Req, Res]{
		Handler: handler,
		Method:  method,
		Path:    path,
		Options: opts,
	}
	h.addRoute(r)
}

func (h *HandlerWithTypes[Req, Res]) addRoute(r *Router) {
	route := r.Echo.Add(h.Method, h.Path, h.Handler, middlewares.Bind[Req])
	echoSwagger.AddRoute[Req, Res](r.OpenAPI, route, h.Options...)
}
