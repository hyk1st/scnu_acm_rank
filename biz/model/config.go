package model

type Config struct {
	VjUserName    string `gorm:"column:vjUserName;type:varchar(255);NOT NULL" json:"vjUserName"`
	VjPassWord    string `gorm:"column:vjPassWord;type:varchar(255);NOT NULL" json:"vjPassWord"`
	VjCookie      string `gorm:"column:vjCookie;type:varchar(2550)" json:"vjCookie"`
	EmailFrom     string `gorm:"column:emailFrom;type:varchar(255);NOT NULL" json:"emailFrom"`
	EmailPassword string `gorm:"column:emailPassword;type:varchar(255);NOT NULL" json:"emailPassword"`
	EmailHost     string `gorm:"column:emailHost;type:varchar(255);NOT NULL" json:"emailHost"`
	EmailSubject  string `gorm:"column:emailSubject;type:text;NOT NULL" json:"emailSubject"`
}

func (m *Config) TableName() string {
	return "config"
}
