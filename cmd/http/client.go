package http

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/client"
	"os"
)

func Client() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: http client <host> <port>")
		os.Exit(1)
	}
	method, path, body := args[0], args[1], args[2]
	myClient := &client.HttpClient{
		Host: "127.0.0.1",
		Port: "8080",
	}
	myClient.Send(method, path, body)
}
