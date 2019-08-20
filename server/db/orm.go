package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	database *xorm.Engine
)

func initOrm() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.args"))

	// create database instance
	if db, err := xorm.NewEngine("mysql", dsn); err != nil {
		panic(err)
	} else {
		database = db
	}

	// options
	database.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	database.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	database.ShowSQL(viper.GetBool("database.show_sql"))

	logrus.Info("connect database success")

	// sync
	return database.StoreEngine("InnoDB").Sync2(
		new(Player),
	)
}
