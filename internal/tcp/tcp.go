package tcp

import (
	"fmt"
	"github.com/mkashifaslam/web-server/internal/http"
	"net"
	"strings"
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
	request := string(incoming[:read])
	fmt.Println(request)

	reqMsg := strings.Split(request, "\n")
	reqStr := strings.Join(reqMsg[2:], "")
	fmt.Println(reqStr)

	// outgoing request
	resStr := fmt.Sprintf("{\"message\":\"%s\"}", reqStr)
	response := http.FormatResponse(resStr, 200, "OK", []http.Header{
		{"Content-Type": "application/json"},
	})

	l, err := conn.Write(response.Format())
	if err != nil {
		panic(err)
	}

	fmt.Println("write bytes", l)
}

func (t *TCP) OpenConn() net.Conn {
	conn, err := net.Dial("tcp", t.Host+":"+t.Port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New connection from %s\n\n", conn.RemoteAddr())
	return conn
}

func (t *TCP) Send(conn net.Conn, method, path, body string) {
	defer conn.Close()

	fmt.Printf("New connection to %s\n\n", conn.RemoteAddr())
	htp := http.FormatRequest(method, path, body, []http.Header{})
	conn.Write(htp.Format())
	response := make([]byte, 1024)
	res, err := conn.Read(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("read bytes", res)
	fmt.Println(string(response[:res]))
}
