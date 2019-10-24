package utils

import "os"

var (
	Secret              string
	PageSize            uint32 = 20
	DefaultExchangeName        = "rabbit.oj"
	CaseQueueName              = "case"
	CaseRoutingKey             = "case"

	JudgeQueueName  = "judge"
	JudgeRoutingKey = "judge"
)

func InitConstant() {
	Secret = os.Getenv("Secret")
}
