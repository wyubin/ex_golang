package main

import (
	"example/swagger_svc/servers"
	"fmt"
	"log"
	"os"
)

const PORT = "4000"

// @title           Swagger Example API
// @version         0.0.1
// @description     This is a sample server for build swagger ui from swaggo yaml.

// @contact.name   YuBin Wang
// @contact.email  yubin.wang@ailabs.tw

// @BasePath  /
func main() {
	errc := make(chan error)

	s := servers.HttpEx{}
	log.Printf("Run server on %s: %+v\n", PORT, s.Info())
	go s.Start(fmt.Sprintf("0.0.0.0:%s", PORT), errc)

	// keep server live
	for err := range errc {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
