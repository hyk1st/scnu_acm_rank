package model

type Competition struct {
	Id          int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"column:name;type:varchar(120);NOT NULL" json:"name"`
	CpId        string `gorm:"column:cp_id;type:varchar(120);NOT NULL" json:"cp_id"`
	Kind        int    `gorm:"column:kind;type:tinyint(4);NOT NULL" json:"kind"`
	StartDate   int64  `gorm:"column:start_date;type:bigint(20);NOT NULL" json:"start_date"`
	Length      int64  `gorm:"column:length;type:bigint(20);NOT NULL" json:"length"`
	Result      string `gorm:"column:result;type:text" json:"result"`
	BestPenalty int    `gorm:"column:bestPenalty;type:int(11);NOT NULL" json:"bestPenalty"`
	BestSolve   int    `gorm:"column:bestSolve;type:int(11);NOT NULL" json:"bestSolve"`
	Ext         string `gorm:"column:ext;type:text;NOT NULL" json:"ext"`
}

func (m *Competition) TableName() string {
	return "competition"
}
