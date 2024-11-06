package servers

type Server interface {
	// start server with addr(run by goroutine)
	Start(addr string, errc chan error)
	Info() map[string]interface{}
}
