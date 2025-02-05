package tcp

import (
	"fmt"
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
	body := fmt.Sprintf("{\"message\":\"Ok!\"}")
	fmt.Println("Content Length:", len(body))
	outgoing := []byte(fmt.Sprintf("HTTP/1.1 200 OK\nContent-Type: application/json\nContent-Length: %d\n\n%s\n", len(body), body))
	l, err := conn.Write(outgoing)
	if err != nil {
		panic(err)
	}

	fmt.Println("write bytes", l)
}
