package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	cfg "ppap/backup/go/config"
)

var DB *sqlx.DB

// Setup 初始化数据库
func Setup() {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.Get("mysql.user"), cfg.Get("mysql.pw"),
		cfg.Get("mysql.host"), cfg.Get("mysql.port"),
		cfg.Get("mysql.db"),
	)
	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil{
		panic(fmt.Sprintf("connect to mysql failed, err:%s", err))
	}
	err = db.Ping()
	if err != nil {
		panic("failed to ping db")
	}

	DB = db
	fmt.Println("init db success")
}