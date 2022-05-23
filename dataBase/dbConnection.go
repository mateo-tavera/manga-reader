package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//TODO: take out connection parameters from the code
const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "my-secret-pw"
	dbname   = "db_mangas"
)

func GetDBConnection() (*sql.DB, error) {
	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	return sql.Open("mysql", mysqlConnection)
}
