package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"fmt"
)

func Register(uid, cid uint32) error {
	contestUser := models.ContestUser{
		Cid: cid,
		Uid: uid,
	}

	if _, err := db.DB.Table("contest_user").Insert(&contestUser);
		err != nil {
		return err
	}

	go UpdateContestParticipant(cid, 1)
	return nil
}

func IsRegistered(uid, cid uint32) (bool, error) {
	contestUser := models.ContestUser{}

	found, err := db.DB.Table("contest_user").
		Select("1").
		Where("uid = ? AND cid = ?", uid, cid).
		Get(&contestUser)

	if err != nil {
		return false, err
	}
	return found, nil
}

func Unregister(uid, cid uint32) error {
	info, err := Info(cid)
	if err != nil {
		return err
	}

	if info.Status != RoundWaiting {
		return errors.New("cannot unregister after the start time of the contest")
	}

	if _, err := db.DB.
		Where("uid = ? AND cid = ?", uid, cid).
		Delete(&models.ContestUser{}); err != nil {
		return err
	}

	go UpdateContestParticipant(cid, -1)
	return nil
}

func UpdateContestParticipant(cid uint32, delta int) {
	if _, err := db.DB.Exec("UPDATE contest SET participants = participants + ? WHERE cid = ?", delta, cid); err != nil {

		fmt.Println(err)
	}
}
