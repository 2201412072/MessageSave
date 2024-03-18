package modelview

type PublicKeyUserSend struct {
	Public_key   string `json:"public_key"`
	Connect_user string `json:"connect_user"`
}

type PublicKeyUserRcv struct {
	Username string `json:"connect_user"`
}
