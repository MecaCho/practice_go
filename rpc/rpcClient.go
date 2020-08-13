package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

func RPCClient() {

	conn, err := net.DialTimeout("tcp", "127.0.0.1:9090", time.Second*10)
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClient(conn)
	fmt.Println("rpc client.")

	var reply string
	err = client.Call("HelloService.HelloWorld", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
