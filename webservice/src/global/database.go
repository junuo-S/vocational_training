package global

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var JunuoDb *sql.DB = nil

func initDataBaseSettings() {
	var err error
	JunuoDb, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err = JunuoDb.Ping(); err != nil {
		panic(err)
	}
	JunuoDb.SetMaxOpenConns(100)
	JunuoDb.SetMaxIdleConns(20)
}
