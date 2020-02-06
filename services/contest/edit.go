package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"time"
)

func Edit(cid string, form *forms.ContestEditForm) error {
	contestInfo := models.Contest{
		Name:      form.Name,
		Status:    form.Status,
		Penalty:   form.Penalty,
		StartTime: models.JSONTime(time.Unix(form.StartTime, 0)),
		EndTime:   models.JSONTime(time.Unix(form.EndTime, 0)),
		BlockTime: models.JSONTime(time.Unix(form.BlockTime, 0)),
	}

	if _, err := db.DB.Table("contest").
		Where("cid = ?", cid).
		Cols("name", "status", "penalty", "start_time", "end_time", "block_time").
		Update(&contestInfo); err != nil {
		return err
	}

	return nil
}
