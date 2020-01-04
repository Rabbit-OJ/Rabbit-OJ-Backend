package contest

import (
	"Rabbit-OJ-Backend/models"
	"time"
)

func CalculateTime(info *models.Contest) int64 {
	start, curr := time.Time(info.StartTime).Unix(), time.Now().Unix()
	return (curr - start) / int64(time.Second)
}
