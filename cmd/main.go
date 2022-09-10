package main

import (
	"log"
	"os"

	"github.com/nachonievag/ip_proxy_api/api"
)

const defaultPort = "8080"

// @title           IP Proxy API
// @version         1.0
// @description     API Providing useful information from ip 2 proxy db

// @host      localhost:8080
func main() {
	log.Println("stating API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
