package routes

import (
	"echotonic/middlewares"
	getuser "echotonic/routes/get_user"
	"net/http"

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

func (r *Router) RegisterRoutes() {
	addRoute[getuser.Request, getuser.Response](r, http.MethodGet, "/users/:id", getuser.Handler, docs.OperationObject{OperationId: "getUserByID"})

}

func addRoute[D any, R any](r *Router, method string, path string, handler echo.HandlerFunc, opts ...docs.OperationObject) {
	route := r.Echo.Add(method, path, handler, middlewares.Bind[D])
	echoSwagger.AddRoute[D, R](r.OpenAPI, route, opts...)
}
