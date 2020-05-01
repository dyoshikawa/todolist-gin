package mysql_util

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func initDb() (*gorp.DbMap, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "secet", "localhost:3307", "dev")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}, nil
}
