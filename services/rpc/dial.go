package rpc

import (
	"fmt"
	"net/rpc"
	"os"
)

type AnyType interface{}

var (
	Client *rpc.Client
)

func DialInit() error {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s", os.Getenv("CASE_DIAL")))
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
