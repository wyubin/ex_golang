package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	fs := http.FileServer(http.Dir(filepath.Join(wd, "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :3000...")
	//err := http.ListenAndServe(":3000", nil)
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
