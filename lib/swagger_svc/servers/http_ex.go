package servers

import (
	"example/swagger_svc/servers/routers"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type HttpEx struct{}

func (s *HttpEx) Start(addr string, errc chan error) {
	var err error
	r := http.NewServeMux()
	// add static server
	wd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("current pwd: %s\n", wd)
	fs := http.FileServer(http.Dir(filepath.Join(wd, "../static")))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	// info route
	rInfo := routers.RouterInfo{}
	rInfo.Init()
	r.Handle("/info/", http.StripPrefix("/info", rInfo.Route()))

	// listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(lis, r)
}

func (s *HttpEx) Info() map[string]interface{} {
	return map[string]interface{}{
		"name": "http-ex-swagger",
	}
}
