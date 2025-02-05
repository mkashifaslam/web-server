package server

import "fmt"

type HttpServer struct {
	Host string
	Port string
}

func (hs *HttpServer) Run() {
	fmt.Printf("Http server running on %s:%s\n", hs.Host, hs.Port)
}
