package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func ClarifyList(cid string) ([]models.ContestClarify, error) {
	var list []models.ContestClarify

	if err := db.DB.Table("contest_clarify").
		Where("cid = ?", cid).
		Desc("ccid").Find(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func ClarifyCreate(cid, message string) (uint32, error) {
	clarify := models.ContestClarify{
		Cid:     cid,
		Message: message,
	}

	if _, err := db.DB.Table("contest_clarify").
		Insert(&clarify); err != nil {
		return 0, err
	}

	return clarify.Ccid, nil
}

func ClarifyDelete(ccid string) error {
	_, err := db.DB.Table("contest_clarify").
		Where("ccid = ?", ccid).
		Delete(&models.ContestClarify{})

	return err
}
