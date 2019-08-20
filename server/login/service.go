package login

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"github.com/sirupsen/logrus"
	"server/protocol"
)

type ServiceLogin struct {
	component.Base
}

func (s *ServiceLogin) Register(sess *session.Session, req *protocol.RegisterReq) error {
	logrus.Info("register")
	return nil
}

func (s *ServiceLogin) Login(sess *session.Session, req *protocol.LoginReq) error {
	logrus.Info("login")
	return nil
}
