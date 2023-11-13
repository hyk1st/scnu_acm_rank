package model

import (
	"time"
)

type Competiton struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(120);NOT NULL" json:"name"`
	VjCpId    string    `gorm:"column:vj_cp_id;type:varchar(120);NOT NULL" json:"vj_cp_id"`
	Kind      int       `gorm:"column:kind;type:tinyint(4);NOT NULL" json:"kind"`
	Password  string    `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	StartDate time.Time `gorm:"column:start_date;type:datetime;NOT NULL" json:"start_date"`
	Length    int64     `gorm:"column:length;type:bigint(20);NOT NULL" json:"length"`
	Result    string    `gorm:"column:result;type:text" json:"result"`
	CreateUsr string    `gorm:"column:create_usr;type:varchar(255);NOT NULL" json:"create_usr"`
	Ext       string    `gorm:"column:ext;type:varchar(255);NOT NULL" json:"ext"`
}

func (m *Competiton) TableName() string {
	return "competiton"
}
