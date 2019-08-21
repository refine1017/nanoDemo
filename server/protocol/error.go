package protocol

import "fmt"

type Error struct {
	ErrMsg string `json:"err_msg"`
}

func (e *Error) WithError(err error) {
	e.ErrMsg = fmt.Sprintf("%v", err)
}

func (e *Error) WithErrorMsg(err string) {
	e.ErrMsg = err
}
