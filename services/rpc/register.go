package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

func Register() {
	fmt.Println("[RPC] registering case service")

	if err := rpc.RegisterName("CaseService", new(CaseService)); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	rpc.Accept(listener)
}
