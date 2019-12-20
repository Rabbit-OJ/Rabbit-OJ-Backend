package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func ClarifyList(cid string) ([]models.ContestClarify, error) {
	var list []models.ContestClarify

	if err := db.DB.Table("contest_clarify").
		Where("cid = ?", cid).
		Order("ccid desc").
		Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ClarifyCreate(cid, message string) (string, error) {
	clarify := models.ContestClarify{
		Cid:     cid,
		Message: message,
	}

	if err := db.DB.Create(&clarify).Error; err != nil {
		return "", err
	}

	return clarify.Ccid, nil
}

func ClarifyDelete(ccid string) error {
	return db.DB.Where("ccid = ?", ccid).Delete(&models.ContestClarify{}).Error
}
