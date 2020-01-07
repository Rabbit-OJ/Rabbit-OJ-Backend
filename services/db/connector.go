package db

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)
