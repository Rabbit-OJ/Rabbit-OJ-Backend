package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/utils"
	"github.com/golang/protobuf/proto"
	"strconv"
)

func Starter(
	code []byte,
	submission *models.Submission,
	questionJudge *models.QuestionJudge,
	question *models.QuestionDetail,
) error {

	request := &protobuf.JudgeRequest{
		Sid:        submission.Sid,
		Tid:        submission.Tid,
		Version:    strconv.FormatUint(uint64(questionJudge.Version), 10),
		Language:   submission.Language,
		TimeLimit:  question.TimeLimit,
		SpaceLimit: question.SpaceLimit,
		CompMode:   questionJudge.Mode,
		Code:       code,
	}

	pro, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	return mq.Publish(
		utils.DefaultExchangeName,
		utils.JudgeRoutingKey,
		pro)
}
