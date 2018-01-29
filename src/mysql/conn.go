package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"

	//"config"
)

var mysqlConn *sql.DB

func init() {
	var mysqlConnErr error
	//var dataSourceName = fmt.Sprintf(
	//	"%s:%s@tcp(%s:%d)/%s?charset=%s",
	//	config.MysqlServerConfig.User,
	//	config.MysqlServerConfig.PassWord,
	//	config.MysqlServerConfig.Host,
	//	config.MysqlServerConfig.Port,
	//	config.MysqlServerConfig.Database,
	//	"utf8")
	var dataSourceName = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"push",
		"utf8")
		fmt.Println( dataSourceName )
	mysqlConn, mysqlConnErr = sql.Open("mysql", dataSourceName)
	if mysqlConnErr != nil {
		log.Fatal("mysqlConnErr===", mysqlConnErr)
	}
	mysqlConn.SetMaxOpenConns(25)
	mysqlConn.SetMaxIdleConns(20)
}
