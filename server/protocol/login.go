package protocol

type LoginReq struct {
	Username string `json:"username"`
}

type LoginRes struct {
	Player *Player `json:"player"`
}
