package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DB() (*sql.DB, error) {
	DBMS := "mysql"
	USER := os.Getenv("user")
	PASS := os.Getenv("password")
	PROTOCOL := os.Getenv("protocol")
	DBNAME := os.Getenv("dbname")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	conn, err := sql.Open(DBMS, CONNECT)
	defer conn.Close()
	if err != nil {
		fmt.Println("Fail to connect db" + err.Error())
		return nil, err
	}
	// 接続確認
	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connect rds : %s", err.Error())
		return nil, err
	}

	fmt.Println("Success to connect rds")
	return conn, nil
}
