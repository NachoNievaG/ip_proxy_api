package main

import (
	"log"
	"os"

	"github.com/nachonievag/ip_proxy_api/api"
)

const defaultPort = "8080"

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
