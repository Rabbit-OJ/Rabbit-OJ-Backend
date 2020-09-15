package initialize

import (
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/submission"
	"context"
)

func Judger(ctx context.Context) {
	judger.InitJudger(ctx)
	judger.OnJudgeResponse = append(judger.OnJudgeResponse, submission.JudgeResponseCallback)
}
