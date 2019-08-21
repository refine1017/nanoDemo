package protocol

type Player struct {
	Id            int64  `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	LastLoginTime int64  `json:"last_login_time"`
	LastLoginIp   string `json:"last_login_ip"`
}
