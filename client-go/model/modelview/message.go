package modelview

// type Message struct {
// 	App          string `json:"app"`
// 	Connect_user string `json:"connect_user"`
// }

type Message struct {
	SrcUser string `json:"connect_user"`
	KeyWord string `json:"app"`
	Operate string `json:"operate"`
}
