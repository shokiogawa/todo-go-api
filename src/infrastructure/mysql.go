package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

// Mysql 他のファイルで使用するためのstruct
type Mysql struct {
	DB *sqlx.DB
}

func NewMysql() (mysql *Mysql, err error) {
	mysql = new(Mysql)
	db, err := tryConnect()
	if err != nil {
		return
	}
	mysql.DB = db
	return
}

func tryConnect() (db *sqlx.DB, err error) {
	var dsn string
	if os.Getenv("MYSQL_HOST_NAME") != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "localhost", os.Getenv("MYSQL_DATABASE"))
	}
	db, err = sqlx.Connect("mysql", dsn)
	return
}

func (mysql *Mysql) Connect() (db *sqlx.DB, err error) {
	err = mysql.DB.Ping()
	if err != nil {
		// リトライ処理
		db, connerr := tryConnect()
		if connerr != nil {
			return nil, connerr
		}
		mysql.DB = db
	}
	return mysql.DB, nil
}
