package model

type Config struct {
	VjUserName string `gorm:"column:vjUserName;type:varchar(255);NOT NULL" json:"vjUserName"`
	VjPassWord string `gorm:"column:vjPassWord;type:varchar(255);NOT NULL" json:"vjPassWord"`
	VjCookie   string `gorm:"column:vjCookie;type:varchar(2550)" json:"vjCookie"`
}

func (m *Config) TableName() string {
	return "config"
}
