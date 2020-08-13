package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloServicer = interface {
	HelloWorld(request string, reply *string) error
}

func RegisterHelloService(svc HelloServicer) error {
	return rpc.RegisterName(HelloServiceName, svc)

}

type HelloService struct {
}

func (h *HelloService) HelloWorld(request string, reply *string) error {
	*reply = fmt.Sprintf("hello: %s.", request)
	fmt.Println("reply: ", *reply)
	return nil
}

func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", "127.0.0.1:9090") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

func RPCServer() {
	RegisterHelloService(new(HelloService))

	// var l net.Listener
	// l, serverAddr := listenTCP()
	// log.Println("Test RPC server listening on", serverAddr)
	// go rpc.Accept(l)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		fmt.Println("accept.")
		go rpc.ServeConn(conn)
	}
}
