package game

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"server/protocol"
)

type ServiceGame struct {
	component.Base
	group *nano.Group
}

func (s *ServiceGame) RpcPlayerLogin(sess *session.Session, req *protocol.RpcPlayerLoginReq) error {
	if s.group == nil {
		s.group = nano.NewGroup("talk")
	}

	return s.group.Add(sess)
}

func (s *ServiceGame) Talk(sess *session.Session, req *protocol.TalkReq) error {
	return s.group.Broadcast("ServiceGame.TalkNotify", &protocol.TalkNotify{Msg: req.Msg})
}
