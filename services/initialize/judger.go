package initialize

import (
	"Rabbit-OJ-Backend/services/judger"
	"context"
)

func Judger(ctx context.Context) {
	judger.InitJudger(ctx)
}
