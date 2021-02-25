package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}