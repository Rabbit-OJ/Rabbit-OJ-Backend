package rpc

import (
	"net"
	"net/rpc"
)

func Register() {
	if err := rpc.RegisterName("CaseService", new(CaseService)); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rpc.ServeConn(conn)
}
