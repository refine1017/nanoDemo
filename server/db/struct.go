package db

import "server/protocol"

type Player struct {
	Id            int64
	Username      string `xorm:"not null VARCHAR(32) default"`
	Nickname      string `xorm:"not null VARCHAR(32) default"`
	LastLoginTime int64  `xorm:"not null BIGINT(20) default"`
	LastLoginIp   string `xorm:"not null VARCHAR(32) default"`
}

func (p *Player) Protocol() protocol.Player {
	return protocol.Player{
		Id:            p.Id,
		Username:      p.Username,
		Nickname:      p.Nickname,
		LastLoginTime: p.LastLoginTime,
		LastLoginIp:   p.LastLoginIp,
	}
}

func (p *Player) Parse(pro protocol.Player) {
	p.Id = pro.Id
	p.Username = pro.Username
	p.Nickname = pro.Nickname
	p.LastLoginTime = pro.LastLoginTime
	p.LastLoginIp = pro.LastLoginIp
}
