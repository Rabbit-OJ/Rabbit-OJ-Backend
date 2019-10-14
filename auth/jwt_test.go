package auth

import (
	"Rabbit-OJ-Backend/models"
	"fmt"
	"testing"
)

func TestSignJWT(t *testing.T) {
	testUser := &models.User{
		Username: "hzytql",
		Uid:      "1",
	}
	result, err := SignJWT(testUser)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(result)
	}
}

func TestVerifyJWT(t *testing.T) {
	oneToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwidXNlcm5hbWUiOiJoenl0cWwiLCJleHAiOjE1NzA5NzA0NzV9.Py-yeOpGebM0H_0ydrAdX0oZ21S4SYGB_nCnKHToEVsGJXG_FeUfJ3m_VnIc_efCvryCu9MKxxp94WASissd_A"

	claims, err := VerifyJWT(oneToken)
	if err != nil {
		t.Fail()
		fmt.Println(err)
	} else {
		fmt.Println(claims.Username)
	}
}
