package routers

import (
	"context"
	"example/swagger_svc/servers/entity"
	"example/swagger_svc/servers/middleware"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type RouterInfo struct {
	route *http.ServeMux
}

func (s *RouterInfo) Init(mws ...endpoint.Middleware) error {
	r := http.NewServeMux()

	r.HandleFunc("GET /{$}", info)
	r.Handle("POST /user", handlerUser(mws...))

	s.route = r
	return nil
}

func (s *RouterInfo) Route() *http.ServeMux {
	return s.route
}

// info godoc
// @Summary      Print info
// @Description  get route info
// @Tags         info
// @Success      200  {string}  string "server info"
// @Router       /info/ [get]
func info(w http.ResponseWriter, r *http.Request) {
	fmt.Print("call info\n")
	fmt.Fprint(w, "server name: swagger test\n")
}

// implement user endpoint
func endpointUser(_ context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*entity.ReqUser)
	resp := map[string]string{
		"user": req.Name,
		"msg":  fmt.Sprintf("hello %s[%d]", req.Name, req.Id),
	}
	return &resp, nil
}

// handlerUser godoc
// @Summary      Say hello to a user
// @Description  get string by ID
// @Tags         info
// @Accept       json
// @Produce      json
// @Param request body entity.ReqUser true "user info"
// @Router       /info/user [post]
func handlerUser(mws ...endpoint.Middleware) http.Handler {
	server := httptransport.NewServer(
		middleware.AggregateMW(mws...)(endpointUser),
		entity.JsonDecoder(&entity.ReqUser{}),
		entity.JsonEncoder(),
	)
	return server
}
