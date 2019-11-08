package utils

import "os"

var (
	Secret              string
	PageSize            uint32 = 20
	DefaultExchangeName        = "rabbit.oj"

	JudgeQueueName  = "judge"
	JudgeRoutingKey = "judge"
	JudgeResultQueueName  = "judge_result"
	JudgeResultRoutingKey = "judge_result"
)

func InitConstant() {
	Secret = os.Getenv("Secret")
}
