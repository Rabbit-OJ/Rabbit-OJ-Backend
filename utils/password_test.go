package utils

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"testing"
)

func TestSaltPassword(t *testing.T) {
	weakPassword := "P@ssw0rd"

	config.InitConstant()
	fmt.Println(config.Secret)

	if salted := SaltPassword(weakPassword); salted != "" {
		fmt.Println(salted)
	} else {
		t.Fail()
	}
}
