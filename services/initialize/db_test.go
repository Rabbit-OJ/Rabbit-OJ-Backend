package initialize

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	Config()
	DB(make(chan bool))

	someMinutesBefore := time.
		Now().
		Add(-1 * 10 * time.Minute)

	var timeoutSubmissions []models.TimeoutSubmission
	if err := db.DB.Table("submission").
		Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
		Updates(map[string]string{"status": "NO", "judge": "[]"}).
		Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(timeoutSubmissions)
}
