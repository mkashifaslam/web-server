package cmd

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/server"
)

func MyHttpServer() {
	fmt.Println("My web server!")
	server.HttpServer()
}
