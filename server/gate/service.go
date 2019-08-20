package gate

import (
	"fmt"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"server/protocol"
)

type ServiceGate struct {
	component.Base
}

func (s *ServiceGate) Hello(sess *session.Session, req *protocol.HelloReq) error {
	name := req.Name
	msg := fmt.Sprintf("hello %v", name)

	return sess.Response(&protocol.HelloRes{Msg: msg})
}
