package utils

import (
	"fmt"
	"testing"
)

func TestSaltPassword(t *testing.T) {
	weakPassword := "P@ssw0rd"

	InitConstant()
	fmt.Println(Secret)

	if salted := SaltPassword(weakPassword); salted != "" {
		fmt.Println(salted)
	} else {
		t.Fail()
	}
}
