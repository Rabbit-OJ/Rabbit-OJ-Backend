package config

import (
	"fmt"
	"path"
	"testing"
)

func TestInitLanguage(t *testing.T) {
	str := "/home/www/111/test.cpp"
	fmt.Println(path.Base(str))
	fmt.Println(path.Dir(str))
}
