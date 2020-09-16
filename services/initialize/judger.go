package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/storage"
	"Rabbit-OJ-Backend/services/submission"
	"context"
)

func Judger(ctx context.Context) {
	judger.InitJudger(ctx, &config.Global.Judger, storage.GetTestCase)
	judger.OnJudgeResponse = append(judger.OnJudgeResponse, submission.JudgeResponseCallback)
}
