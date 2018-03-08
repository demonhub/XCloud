package db

import (
	"config"
	"github.com/jmoiron/sqlx"
)

var DB *Database

type Database struct {
	Conn 	*sqlx.DB
	State 	bool
}

func InitDatabase() {
	Conn,err := sqlx.Open(config.DatabaseDriver, config.DatabaseUSER + ":" + config.DatabasePWD + "@" + config.DatabaseURL)
	checkErr(err)
	DB = &Database{Conn:Conn,State:true}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}