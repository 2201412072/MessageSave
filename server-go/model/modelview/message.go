package modelview

type Message struct {
	SrcUser string `gorm:"primaryKey;column:src_user"`
	DstUser string `gorm:"primaryKey;column:dst_user"`
	KeyWord string `gorm:"primaryKey;size:100"`
	Operate string `gorm:"size:100"`
}
