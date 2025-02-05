package server

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/tcp"
)

type HttpServer struct {
	Host string
	Port string
}

func (hs *HttpServer) Run() {
	transport := tcp.New(hs.Host, hs.Port)
	fmt.Printf("Http server running on http://%s:%s\n", hs.Host, hs.Port)
	transport.ListenAndAccept()
}
