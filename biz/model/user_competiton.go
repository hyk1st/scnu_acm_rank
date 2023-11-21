package model

import (
	"fmt"
	"time"
)

type UserCompetition struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VjName  string `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CompId  int    `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	Solve   int    `gorm:"column:solve;type:int(11);NOT NULL" json:"solve"`
	Rank    uint   `gorm:"column:rank;type:tinyint(4) unsigned;NOT NULL" json:"rank"`
	Ext     string `gorm:"column:ext;type:text" json:"ext"`
	Penalty int    `gorm:"column:penalty;type:int(11);NOT NULL" json:"penalty"`
}

func (m *UserCompetition) TableName() string {
	return "user_competition"
}

type Result struct {
	Name      string `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`
	StuId     int64  `gorm:"column:stu_id;type:bigint(20);NOT NULL" json:"stu_id"`
	VjName    string `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CompId    int    `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	CompName  string `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	Penalty   int    `gorm:"column:penalty;type:int(11)" json:"penalty"`
	Rank      int    `gorm:"column:rank;type:int(11);NOT NULL" json:"rank"`
	Solve     int    `gorm:"column:solve;type:int(11);NOT NULL" json:"solve"`
	StartDate int64  `gorm:"column:start_date;type:bigint(20);NOT NULL" json:"start_date"`
}

func GetUserCompetitions() ([]Result, error) {
	time := time.Now().Add(-time.Hour * 24 * 180).Unix()
	sql := `SELECT user.name, user.stu_id, user.vj_name, a.group_id, b.name comp_name, a.rank, b.start_date 
			FROM user, user_competition a, competiton b
			WHERE user.vj_name = a.vj_name AND a.comp_id = b.id AND b.start_date > 0 AND b.kind < 2 
			ORDER BY user.name, a.rank DESC`
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
