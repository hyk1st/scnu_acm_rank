package model

import (
	"time"
)

type UserCompetition struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VjName  string `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CompId  int    `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	Solve   int    `gorm:"column:solve;type:int(11);NOT NULL" json:"solve"`
	Rank    uint   `gorm:"column:rank;type:tinyint(4) unsigned;NOT NULL" json:"rank"`
	Score   int    `gorm:"column:score;type:int(11)" json:"score"`
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
	Score     int    `gorm:"column:score;type:int(11)" json:"score"`
	Solve     int    `gorm:"column:solve;type:int(11);NOT NULL" json:"solve"`
	StartDate int64  `gorm:"column:start_date;type:bigint(20);NOT NULL" json:"start_date"`
}

type GroupResult struct {
	StuId   int64  `gorm:"column:stu_id;type:bigint(20);NOT NULL" json:"stu_id"`
	Name    string `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`
	GroupId int64  `gorm:"column:group_id;type:bigint(20);NOT NULL" json:"group_id"`
	//GroupName    string `gorm:"column:_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CompId    int    `gorm:"column:comp_id;type:int(11);NOT NULL" json:"comp_id"`
	CompName  string `gorm:"column:comp_name;type:varchar(255);NOT NULL" json:"comp_name"`
	Penalty   int    `gorm:"column:penalty;type:int(11)" json:"penalty"`
	Rank      int    `gorm:"column:rank;type:int(11);NOT NULL" json:"rank"`
	Score     int    `gorm:"column:score;type:int(11)" json:"score"`
	Solve     int    `gorm:"column:solve;type:int(11);NOT NULL" json:"solve"`
	StartDate int64  `gorm:"column:start_date;type:bigint(20);NOT NULL" json:"start_date"`
}

func GetUserCompetitions() ([]Result, error) {
	time := time.Now().Add(-time.Hour * 24 * 700).UnixMilli()
	sql := `SELECT user.name, user.stu_id, user.vj_name, a.comp_id, b.name comp_name, a.rank, a.penalty, a.score, a.solve, b.start_date 
			FROM user, user_competition a, competition b
			WHERE (user.vj_name = a.vj_name AND a.comp_id = b.id AND b.start_date > ? AND b.kind = 0) OR (user.nc_name = a.vj_name AND a.comp_id = b.id AND b.start_date > ? AND b.kind = 1)
			ORDER BY user.name, a.score DESC`
	res := make([]Result, 0)
	DB.Raw(sql, time, time).Find(&res)
	return res, nil
}

func GetGroupCompetitions() ([]GroupResult, error) {
	ti := time.Now().Add(-time.Hour * 24 * 180)
	sql := `SELECT rr.group_id, tt.name, tc.name comp_name, tu.rank, tu.solve, tu.penalty, tu.score, tc.start_date FROM
(SELECT user.group_id, a.comp_id, MIN(a.rank) rk
FROM user, user_competition a, competition b
WHERE ((user.vj_name = a.vj_name AND a.comp_id = b.id AND b.kind = 13) OR (user.nc_name = a.vj_name AND a.comp_id = b.id AND b.kind = 14)) AND b.start_date > ? AND group_id > 0
GROUP BY user.group_id, a.comp_id
) rr, team tt, user_competition tu, competition tc
WHERE rr.group_id = tt.id  
AND rr.comp_id = tu.comp_id
AND rr.rk = tu.rank 
AND rr.comp_id = tc.id;`
	res := make([]GroupResult, 0)
	DB.Raw(sql, ti).Find(&res)
	return res, nil
}

func GetGroupCompetitionsByTeam(id string) ([]GroupResult, error) {
	ti := time.Now().Add(-time.Hour * 24 * 180)
	sql := `SELECT rr.group_id, tt.name, tc.name comp_name, tu.rank, tu.solve, tu.penalty, tu.score, tc.start_date FROM
(SELECT user.group_id, a.comp_id, MIN(a.rank) rk
FROM user, user_competition a, competition b
WHERE user.group_id = ? AND ((user.vj_name = a.vj_name AND a.comp_id = b.id AND b.kind = 13) OR (user.nc_name = a.vj_name AND a.comp_id = b.id AND b.kind = 14)) AND b.start_date > ? AND group_id > 0
GROUP BY user.group_id, a.comp_id
) rr, team tt, user_competition tu, competition tc
WHERE rr.group_id = tt.id  
AND rr.comp_id = tu.comp_id
AND rr.rk = tu.rank 
AND rr.comp_id = tc.id;
`
	res := make([]GroupResult, 0)
	DB.Raw(sql, id, ti).Find(&res)
	return res, nil
}
