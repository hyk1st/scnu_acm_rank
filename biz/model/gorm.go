package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runtime"
	"sync"
)

var DB *gorm.DB = nil
var once sync.Once

func GetDB() *gorm.DB {

	dsn := "root:123456@tcp(172.26.144.1:3306)/scnu_acm_rank?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	runtime.KeepAlive(DB)
	return DB
}
