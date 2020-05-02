package mysql_util

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorp.DbMap

func InitDb() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "secet", "localhost:3307", "dev")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	Db = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return nil
}
