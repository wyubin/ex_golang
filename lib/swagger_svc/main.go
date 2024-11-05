package main

import (
	"example/swagger_svc/route"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	r := http.NewServeMux()
	// static server
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	fmt.Printf("current pwd: %s\n", wd)
	fs := http.FileServer(http.Dir(filepath.Join(wd, "static")))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	// other route
	exRoute := route.NewRoute()
	r.Handle("/route/", http.StripPrefix("/route", exRoute))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
