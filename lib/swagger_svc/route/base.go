package route

import (
	"fmt"
	"net/http"
)

func NewRoute() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /home", home)

	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Print("call home\n")
	fmt.Fprint(w, "home\n")
}
