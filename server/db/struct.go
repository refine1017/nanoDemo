package db

type Player struct {
	Id            int64
	Username      string `xorm:"not null VARCHAR(32) default"`
	Nickname      string `xorm:"not null VARCHAR(32) default"`
	LastLoginTime int64  `xorm:"not null BIGINT(20) default"`
	LastLoginIp   string `xorm:"not null VARCHAR(32) default"`
}
