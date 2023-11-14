package model

import (
	"fmt"
	"time"
)

type UserCompetiton struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	UserId    int       `gorm:"column:user_id;type:int(11);NOT NULL" json:"user_id"`
	UserName  string    `gorm:"column:user_name;type:varchar(10);NOT NULL" json:"user_name"`
	CompName  string    `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	CompId    int       `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	Goal      int64     `gorm:"column:goal;type:bigint(20);NOT NULL" json:"goal"`
	StartTime time.Time `gorm:"column:start_time;type:datetime;NOT NULL" json:"start_time"`
	Ext       string    `gorm:"column:ext;type:varchar(255)" json:"ext"`
}

func (m *UserCompetiton) TableName() string {
	return "user_competiton"
}

type Result struct {
	Name      string    `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`
	StuId     int64     `gorm:"column:stu_id;type:bigint(20);NOT NULL" json:"stu_id"`
	VjName    string    `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	GroupId   int       `gorm:"column:group_id;type:int(11);NOT NULL" json:"group_id"`
	CompName  string    `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	Kind      int       `gorm:"column:kind;type:tinyint(4)" json:"kind"`
	Goal      int64     `gorm:"column:goal;type:bigint(20);NOT NULL" json:"goal"`
	StartTime time.Time `gorm:"column:start_time;type:datetime;NOT NULL" json:"start_time"`
}

func GetUserCompetitions() ([]Result, error) {
	time := time.Now().Add(-time.Hour * 24 * 180)
	sql := "SELECT user.name, user.stu_id, user.group_id,user.vj_name, temp.comp_name, temp.goal, temp.start_time " +
		"FROM user, user_competition temp " +
		"WHERE user.name = temp.user_name AND temp.start_time > ? AND temp.kind < 2" +
		"ORDER BY user.name, temp.goal DESC"
	res := make([]Result, 0)
	DB.Raw(sql, time).Find(&res)
	fmt.Println(res)
	return res, nil
}

func GetGroupCompetitions() ([]Result, error) {
	time := time.Now().Add(-time.Hour * 24 * 180)
	sql := "SELECT user.name, user.stu_id, user.group_id, user.vj_name, temp.comp_name, temp.goal, temp.start_time " +
		"FROM user, user_competition temp " +
		"WHERE user.name = temp.user_name AND temp.start_time > ? AND temp.kind >= 2" +
		"ORDER BY user.name, temp.goal DESC"
	res := make([]Result, 0)
	DB.Raw(sql, time).Find(&res)
	fmt.Println(res)
	return res, nil
}
