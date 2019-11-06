package rpc

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/question"
)

type CaseService struct {}

func (s *CaseService) Case(request protobuf.TestCaseRequest, response *protobuf.TestCaseResponse) error {
	tid := request.Tid
	err := question.Case(tid, response)

	if err != nil {
		return err
	}

	return nil
}