package model

type User struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Email      string `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	Password   string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	VjName     string `gorm:"column:vj_name;type:varchar(50);NOT NULL" json:"vj_name"`
	CfId       string `gorm:"column:cf_id;type:varchar(50);NOT NULL" json:"cf_id"`
	StuId      int64  `gorm:"column:stu_id;type:bigint(20);NOT NULL" json:"stu_id"`
	Name       string `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`
	Sex        int    `gorm:"column:sex;type:tinyint(4);NOT NULL" json:"sex"`
	GroupId    int    `gorm:"column:group_id;type:int(11);NOT NULL" json:"group_id"`
	Grade      string `gorm:"column:grade;type:varchar(10);NOT NULL" json:"grade"`
	Status     int    `gorm:"column:status;type:tinyint(4);NOT NULL" json:"status"`
	Level      int    `gorm:"column:level;type:tinyint(4);NOT NULL" json:"level"`
	Connection string `gorm:"column:connection;type:varchar(255)" json:"connection"`
	Ext        string `gorm:"column:ext;type:varchar(255)" json:"ext"`
}

func (m *User) TableName() string {
	return "user"
}
