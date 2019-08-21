package protocol

type LoginReq struct {
	Username string `json:"username"`
}

type LoginRes struct {
	Error
	Player Player `json:"player"`
}
