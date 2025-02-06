package http

import (
	"github.com/mkashifaslam/web-server/internal/server"
)

func Server() {
	myServer := &server.HttpServer{
		Host: "127.0.0.1",
		Port: "8080",
	}
	myServer.Run()
}
