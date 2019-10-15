package question

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Delete(tid string) error {
	tx := db.DB.Begin()

	// TODO: Add delete file logic...

	if err := tx.Where("tid = ?", tid).Delete(models.Question{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("tid = ?", tid).Delete(models.QuestionContent{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("tid = ?", tid).Delete(models.QuestionJudge{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
