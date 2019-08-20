package db

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"server/protocol"
)

var playerManager = NewPlayerManager()

type PlayerManager struct {
	component.Base
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{}
}

func (m *PlayerManager) Login(s *session.Session, req *protocol.LoginReq) error {
	return nil
}
