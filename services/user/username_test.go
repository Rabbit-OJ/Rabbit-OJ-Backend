package user

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"testing"
)

func TestUsernameToUid(t *testing.T) {
	db.Init()

	uid, err := UsernameToUid("root")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(uid)
	}
}
