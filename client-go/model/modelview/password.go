package modelview

// type Password struct {
// 	App          string `gorm:"primaryKey;column:application"`
// 	Connect_user string `gorm:"primaryKey;column:user"`
// 	Password     string `gorm:"size:100"`
// }

type Password struct {
	Application string `json:"app"`
	User        string `json:"connect_user"`
	Password    string `json:"password"`
}
