package utils

import "os"

var (
	Secret string
)

func InitConstant() {
	Secret = os.Getenv("Secret")
}