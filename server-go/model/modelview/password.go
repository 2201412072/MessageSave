package modelview

type Password struct {
	Application string `gorm:"primaryKey;column:application"`
	User        string `gorm:"primaryKey;column:user"`
	Password    string `gorm:"size:100"`
}
