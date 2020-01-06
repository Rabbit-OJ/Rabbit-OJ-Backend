package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)
