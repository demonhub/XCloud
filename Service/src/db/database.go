package db

import (
	"database/sql"
	"fmt"
	"config"
)

var DB *Database

type Database struct {
	Conn *sql.DB
}

func InitDatabase() {
	Conn,err := sql.Open(config.DatabaseDriver, config.DatabaseUSER + ":" + config.DatabasePWD + "@" + config.DatabaseURL)
	checkErr(err)
	DB = &Database{Conn:Conn}
}

func (DB *Database)Exec(query string, args ... interface{}) error {
	stmt, err := DB.Conn.Prepare(query)
	checkErr(err)
	res, err := stmt.Exec(args)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("Exec query -> '",query,"' parm -> ",args)
}

func (DB *Database)Insert(query string, args ... interface{}) error {
	stmt, err := DB.Conn.Prepare(query)
	checkErr(err)
	res, err := stmt.Exec(args)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("Exec query -> '",query,"' parm -> ",args)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}