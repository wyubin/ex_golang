package main

import "fmt"

type Opt struct {
	host    string
	maxConn int
	tls     bool
}

// 用來改變 Opt 值的 func
type OptFunc func(*Opt)

// 要有預設值
func defOpt() Opt {
	return Opt{
		maxConn: 10,
		tls:     false,
	}
}

// service 就直接繼承 Opt 就可以
type Server struct {
	Opt
}

// new, require argument 再加上 optional
func newServer(host string, opts ...OptFunc) *Server {
	o := defOpt()
	// 依序改 opt 後
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opt: o,
	}
}

// add a OptFunc
func withMaxConn(numConn int) OptFunc {
	return func(o *Opt) {
		o.maxConn = numConn
	}
}

func main() {
	s := newServer("1.1.1.1")
	fmt.Printf("default server: %+v\n", s)
	sConn100 := newServer("1.1.1.1", withMaxConn(100))
	fmt.Printf("custom server: %+v\n", sConn100)
}
