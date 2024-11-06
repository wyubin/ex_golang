package routers

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Router interface {
	// init router
	Init(mws ...endpoint.Middleware) error
	// return router handler
	Route() *http.ServeMux
}
