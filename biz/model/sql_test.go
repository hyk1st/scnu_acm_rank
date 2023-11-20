package model

import (
	"fmt"
	"github.com/henrylee2cn/ameda"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGetUserCompetitions(t *testing.T) {
	GetDB()
	_, _ = GetUserCompetitions()
}

func TestSh2db(t *testing.T) {
	dsn := "root:123456@tcp(172.26.144.1:3306)/scnu?charset=utf8mb4&parseTime=True&loc=Local"
	GetDB()
	var err error
	db2, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	list := make([]User2, 0)
	db2.Model(&User2{}).Find(&list)
	for _, v := range list {
		id, _ := ameda.StringToInt64(v.StudentNumber)
		user := &User{
			//Id:         0,
			//Email:      ,
			//Password:   ,
			VjName: v.VjudgeName,
			CfId:   v.CfId,
			StuId:  id,
			Name:   v.TrueName,
		}
		DB.Model(&User{}).Save(&user)
	}
}
