package mysql

import (
	"fmt"
	"go/go/src/error"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type tb struct {
	db        *sqlx.DB
	tableName string
}

type selftb struct {
}

type dbc interface {
}

func intdatabase(databaseName string, selfTable interface{}) *tb {

	table := new(tb)
	database, err := sqlx.Open("mysql", "root:123456@tcp(47.93.125.10:3306)/"+databaseName)
	error.HandleError(err, "open mysql failed")
	fmt.Println(database)
	table.db = database
	return table
}
