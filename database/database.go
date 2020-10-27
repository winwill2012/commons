package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

// 调用该方法需要引入MySQL驱动：_ "github.com/go-sql-driver/mysql"
func ConnectMySQL(hostname string, port int, database, username, password string) (*sqlx.DB, error) {
	connConf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", username, password, hostname, port, database)
	db, err := sqlx.Open("mysql", connConf)
	if err != nil {
		log.Printf("连接数据库%s失败", database)
		return nil, err
	}
	log.Printf("连接数据库%s成功", database)
	return db, nil
}

// 调用该方法需要引入SQLServer驱动：_ "github.com/denisenkom/go-mssqldb"
func ConnectSQLServer(hostname string, port int, database, username, password string) (*sqlx.DB, error) {
	connConf := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=disable",
		hostname, database, username, password)
	db, err := sqlx.Open("mssql", connConf)
	if err != nil {
		log.Printf("连接数据库%s失败", database)
		return nil, err
	}
	log.Printf("连接数据库%s成功", database)
	return db, nil
}
