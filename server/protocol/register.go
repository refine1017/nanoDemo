package protocol

type RegisterReq struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type RegisterRes struct {
	Player *Player `json:"player"`
}
