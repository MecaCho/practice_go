package go_tcp

import (
	"fmt"
	"net"
	"strings"
)

func connhandler(conn net.Conn) {

	if conn == nil {
		fmt.Println("conn is nil..., return")
		return
	}

	buf := make([]byte, 4096)
	for {

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
		inputStr := strings.TrimSpace(string(buf[:cnt]))
		switch {
		case strings.Contains(inputStr, "ping"):
			fmt.Println("receive from client: ", inputStr)
			conn.Write([]byte("pong\n"))
		default:
			fmt.Println(cnt, inputStr)
		}
	}

	fmt.Println("connect closed:", conn.RemoteAddr())

}

func TCPServer() {

	server, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("server accept error")
		}
		go connhandler(conn)
	}
}

func HandlerClient(conn net.Conn) {
	for {
		conn.Write([]byte("hello i am client write"))
		conn.Read([]byte("hello i am client read"))
	}
}

func TCPClient() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		fmt.Println("tcp client dial error", err.Error())
	}

	HandlerClient(conn)
}
