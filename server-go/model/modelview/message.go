package modelview

type Message struct {
	SrcUser string `json:"src_user"`
	DstUser string `json:"dst_user"`
	KeyWord string `json:"keyword"`
	Operate string `json:"operator"`
}
