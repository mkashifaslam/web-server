package client

import (
	"github.com/mkashifaslam/web-server/internal/tcp"
)

type HttpClient struct {
	Host string
	Port string
}

func (c *HttpClient) Send(method, path, body string) {
	transport := tcp.New(c.Host, c.Port)
	conn := transport.OpenConn()
	transport.Send(conn, method, path, body)
}
