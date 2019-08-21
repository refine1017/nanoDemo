package login

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"server/protocol"
)

type ServiceLogin struct {
	component.Base
}

func (s *ServiceLogin) Register(sess *session.Session, req *protocol.RegisterReq) error {
	return sess.RPC("ServiceDB.RpcCreatePlayer", &protocol.RpcCreatePlayerReq{Username: req.Username, Nickname: req.Nickname})
}

func (s *ServiceLogin) Login(sess *session.Session, req *protocol.LoginReq) error {
	return sess.RPC("ServiceDB.RpcLoadPlayer", &protocol.RpcLoadPlayerReq{Username: req.Username})
}
