package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Register(uid, cid string) error {
	contestUser := models.ContestUser{
		Cid: cid,
		Uid: uid,
	}

	if err := db.DB.Create(&contestUser).Error; err != nil {
		return err
	}

	go UpdateContestParticipant(cid, 1)
	return nil
}

func IsRegistered(uid, cid string) (bool, error) {
	count := 0

	if err := db.DB.Table("contest_user").
		Select("1").
		Where("uid = ? AND cid = ?", uid, cid).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count == 1, nil
}

func Unregister(uid, cid string) error {
	if err := db.DB.
		Where("uid = ? AND cid = ?", uid, cid).
		Delete(&models.ContestUser{}).
		Error; err != nil {
		return err
	}

	go UpdateContestParticipant(cid, -1)
	return nil
}

func UpdateContestParticipant(cid string, delta int) {
	if err := db.DB.Table("contest").
		Where("cid = ?", cid).
		Update("participants", gorm.Expr("participants + ?", delta)).
		Error;
		err != nil {

		fmt.Println(err)
	}
}
