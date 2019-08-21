package protocol

type TalkReq struct {
	Msg string `json:"msg"`
}

type TalkNotify struct {
	Msg string `json:"msg"`
}
