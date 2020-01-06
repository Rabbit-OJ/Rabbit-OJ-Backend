package contest

import (
	"Rabbit-OJ-Backend/models"
	"time"
)

func CalculateTime(info *models.Contest) uint32 {
	start, curr := time.Time(info.StartTime).Unix(), time.Now().Unix()
	return uint32(curr - start)
}
