package conn

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func Connect() {
	DB, err = sql.Open("postgres", "dbname=bongobitan host=localhost sslmode=disable")
	dieIf(err)
	err = DB.Ping()
	dieIf(err)
}
func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
