package model

import (
	"fmt"
	"time"
)

type UserCompetition struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VjName    string `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CompName  string `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	CompId    int    `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	Rank      uint   `gorm:"column:rank;type:tinyint(4) unsigned;NOT NULL" json:"rank"`
	Goal      int64  `gorm:"column:goal;type:bigint(20);NOT NULL" json:"goal"`
	Kind      int    `gorm:"column:kind;type:tinyint(4)" json:"kind"`
	StartTime int64  `gorm:"column:start_time;type:bigint(20);NOT NULL" json:"start_time"`
	Ext       string `gorm:"column:ext;type:text" json:"ext"`
}

func (m *UserCompetition) TableName() string {
	return "user_competition"
}

type Result struct {
	Name      string    `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`
	StuId     int64     `gorm:"column:stu_id;type:bigint(20);NOT NULL" json:"stu_id"`
	VjName    string    `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	GroupId   int       `gorm:"column:group_id;type:int(11);NOT NULL" json:"group_id"`
	CompName  string    `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	Kind      int       `gorm:"column:kind;type:tinyint(4)" json:"kind"`
	Rank      int       `gorm:"column:rank;type:int(11);NOT NULL" json:"rank"`
	StartTime time.Time `gorm:"column:start_time;type:datetime;NOT NULL" json:"start_time"`
}

func GetUserCompetitions() ([]Result, error) {
	time := time.Now().Add(-time.Hour * 24 * 180)
	sql := "SELECT user.name, user.stu_id, temp.comp_name, temp.goal, temp.start_time " +
		"FROM user, user_competition temp " +
		"WHERE user.vj_name = temp.vj_name AND temp.start_time > ? AND temp.kind < 2" +
		"ORDER BY user.name, temp.goal DESC"
	res := make([]Result, 0)
	DB.Raw(sql, time).Find(&res)
	fmt.Println(res)
	return res, nil
}

func GetGroupCompetitions() ([]Result, error) {
	ti := time.Now().Add(-time.Hour * 24 * 180)
	sql := "SELECT temp.name, temp.stu_id, temp.comp_name, temp.goal, temp.start_time " +
		"FROM user_competition temp " +
		"WHERE temp.start_time > ? AND temp.kind >= 2" +
		"ORDER BY temp.stu_id, temp.goal DESC"
	res := make([]Result, 0)
	DB.Raw(sql, ti).Find(&res)
	fmt.Println(res)
	return res, nil
}
