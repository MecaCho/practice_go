package rpc

import (
	"testing"
)

// func startServer() {
// 	Register(new(Arith))
// 	Register(new(Embed))
// 	RegisterName("net.rpc.Arith", new(Arith))
// 	Register(BuiltinTypes{})
//
// 	var l net.Listener
// 	l, serverAddr = listenTCP()
// 	log.Println("Test RPC server listening on", serverAddr)
// 	go Accept(l)
//
// 	HandleHTTP()
// 	httpOnce.Do(startHttpServer)
// }

func TestRPCServer(t *testing.T) {
	RPCServer()
}

func TestRPCClient(t *testing.T) {
	RPCClient()
}
