package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/storage"
	"Rabbit-OJ-Backend/services/submission"
	"context"
	"github.com/Rabbit-OJ/Rabbit-OJ-Judger"
	"os"
)

func Judger(ctx context.Context) {
	role := os.Getenv("Role")

	if role == "Judge" {
		judger.InitJudger(ctx, &config.Global.Judger, storage.GetTestCase, true, true, role)
	} else if role == "Server" {
		judger.InitJudger(ctx, &config.Global.Judger, storage.GetTestCase, false, true, role)
	}
	judger.OnJudgeResponse = append(judger.OnJudgeResponse, submission.JudgeResponseCallback)
}
