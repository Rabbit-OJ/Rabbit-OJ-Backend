package utils

import "os"

var (
	Secret string
	PageSize uint32 = 20
)

func InitConstant() {
	Secret = os.Getenv("Secret")
}