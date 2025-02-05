package cmd

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/server"
)

func MyServer() {
	fmt.Println("My web server!")
	server.HttpServer()
}
