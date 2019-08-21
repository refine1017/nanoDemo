package db

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"server/protocol"
)

type ServiceDB struct {
	component.Base
}

func (s *ServiceDB) RpcCreatePlayer(sess *session.Session, req *protocol.RpcCreatePlayerReq) error {
	exists, err := database.Exist(&Player{
		Username: req.Username,
	})

	res := &protocol.RegisterRes{}

	if err != nil {
		res.WithError(err)
		return sess.Push("ServiceLogin.Register", res)
	}

	if exists {
		res.WithErrorMsg("account exists")
		return sess.Push("ServiceLogin.Register", res)
	}

	player := &Player{}

	if _, err := database.InsertOne(&Player{
		Username: req.Username,
		Nickname: req.Nickname,
	}); err != nil {
		res.WithError(err)
		return sess.Push("ServiceLogin.Register", res)
	}

	has, err := database.Table(&Player{}).Where("username = ?", req.Username).Get(player)
	if err != nil {
		res.WithError(err)
		return sess.Push("ServiceLogin.Register", res)
	}

	if !has {
		res.WithErrorMsg("no found")
		return sess.Push("ServiceLogin.Register", res)
	}

	res.Player = player.Protocol()
	return sess.Push("ServiceLogin.Register", res)
}

func (s *ServiceDB) RpcLoadPlayer(sess *session.Session, req *protocol.RpcLoadPlayerReq) error {
	player := &Player{}

	res := &protocol.LoginRes{}

	has, err := database.Table(&Player{}).Where("username = ?", req.Username).Get(player)
	if err != nil {
		res.WithError(err)
		return sess.Push("ServiceLogin.Register", res)
	}

	if !has {
		res.WithErrorMsg("no found")
		return sess.Push("ServiceLogin.Register", res)
	}

	res.Player = player.Protocol()
	return sess.Push("ServiceLogin.Login", res)
}

func (s *ServiceDB) RpcSavePlayer(sess *session.Session, req *protocol.RpcSavePlayerReq) error {
	player := &Player{}
	player.Parse(req.Player)

	_, err := database.Update(player, &Player{Username: player.Username})
	if err != nil {
		return err
	}

	return nil
}
