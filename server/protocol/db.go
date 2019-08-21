package protocol

type RpcCreatePlayerReq struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type RpcLoadPlayerReq struct {
	Username string `json:"username"`
}

type RpcSavePlayerReq struct {
	Player Player `json:"player"`
}
