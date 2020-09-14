package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/mq"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
)

func Starter(
	code []byte,
	submission *models.Submission,
	questionJudge *models.QuestionJudge,
	question *models.QuestionDetail,
	isContest bool,
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
		Time:       time.Now().Unix(),
		IsContest:  isContest,
	}

	pro, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	return mq.PublishMessage(
		config.JudgeRequestTopicName,
		[]byte(fmt.Sprintf("%d%d", submission.Sid, submission.Tid)),
		pro)
}
