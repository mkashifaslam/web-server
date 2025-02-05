package cmd

import (
	"github.com/mkashifaslam/web-server/internal/server"
)

func MyHttpServer() {
	myServer := &server.HttpServer{
		Host: "127.0.0.1",
		Port: "8080",
	}
	myServer.Run()
}
