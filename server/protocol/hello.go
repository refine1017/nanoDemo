package protocol

type HelloReq struct {
	Name string `json:"name"`
}

type HelloRes struct {
	Msg string `json:"msg"`
}
