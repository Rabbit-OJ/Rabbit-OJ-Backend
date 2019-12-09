package rpc

import (
	"Rabbit-OJ-Backend/services/config"
	"net/rpc"
)

type AnyType interface{}

var (
	Client *rpc.Client
)

func DialInit() error {
	client, err := rpc.Dial("tcp", config.Global.Rpc)
	if err != nil {
		return err
	}

	Client = client

	return nil
}

func DialCall(serviceName, functionName string, request AnyType, response AnyType) error {
	if Client == nil {
		if err := DialInit(); err != nil {
			return err
		}
	}

	if err := Client.Call(serviceName+"."+functionName, request, response); err != nil {
		return err
	}

	return nil
}
