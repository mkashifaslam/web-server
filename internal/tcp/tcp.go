package tcp

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/http"
	"net"
)

type TCP struct {
	Host string
	Port string
}

func New(host, port string) *TCP {
	return &TCP{
		Host: host,
		Port: port,
	}
}

func (t *TCP) ListenAndAccept() net.Listener {
	ls, err := net.Listen("tcp", t.Host+":"+t.Port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}

	return ls
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("New connection from %s\n\n", conn.RemoteAddr())

	// incoming request
	incoming := make([]byte, 1024)
	read, err := conn.Read(incoming)
	if err != nil {
		return
	}

	fmt.Println("read bytes", read)
	fmt.Println(string(incoming[:read]))

	// outgoing request
	response := http.Response("{\"message\":\"Ok!\"}", 200, "OK", []http.Header{
		{"Content-Type": "application/json"},
	})

	l, err := conn.Write(response.Format())
	if err != nil {
		panic(err)
	}

	fmt.Println("write bytes", l)
}
