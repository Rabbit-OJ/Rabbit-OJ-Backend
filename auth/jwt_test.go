package auth

import (
	"Rabbit-OJ-Backend/models"
	"fmt"
	"testing"
)

func TestSignJWT(test *testing.T) {
	testUser := &models.User{
		Username: "hzytql",
		Uid:      "1",
	}
	result, err := SignJWT(testUser)

	if err != nil {
		fmt.Println(err)
		test.Fail()
	} else {
		fmt.Println(result)
	}
}
