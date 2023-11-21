package model

type Team struct {
	Id     int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Name   string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Key    string `gorm:"column:key;type:varchar(255);NOT NULL" json:"key"`
	Leader int64  `gorm:"column:leader;type:bigint(20);NOT NULL" json:"leader"`
	Status int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`
	Ext    string `gorm:"column:ext;type:varchar(255)" json:"ext"`
}

func (m *Team) TableName() string {
	return "team"
}
