package goroutine

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func TcpTimer() {
	listener, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func NetCat() {

	conn, err := net.Dial("tcp", "localhost:8003")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	MustCopy(os.Stdout, conn)
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
