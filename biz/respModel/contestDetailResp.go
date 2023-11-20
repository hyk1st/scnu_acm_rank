package respModel

import (
	"scnu_acm_rank/biz/remote"
)

type ContestDetailResp struct {
	Name      string `gorm:"column:name;type:varchar(120);NOT NULL" json:"name"`
	VjCpId    string `gorm:"column:vj_cp_id;type:varchar(120);NOT NULL" json:"vj_cp_id"`
	Kind      int    `gorm:"column:kind;type:tinyint(4);NOT NULL" json:"kind"`
	Password  string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	StartDate int64  `gorm:"column:start_date;type:bigint(20);NOT NULL" json:"start_date"`
	Length    int64  `gorm:"column:length;type:bigint(20);NOT NULL" json:"length"`
	CreateUsr string `gorm:"column:create_usr;type:varchar(255);NOT NULL" json:"create_usr"`
	Result    remote.AnalysisRes
}
