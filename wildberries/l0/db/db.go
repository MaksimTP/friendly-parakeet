package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "wb_lvl0"
)

type DataBase struct {
	db *sql.DB
}

var DbInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func (d *DataBase) Connect(driverName string, dbInfo string) {
	db, err := sql.Open(driverName, dbInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	d.db = db
}

func (d *DataBase) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

// func main() {
// 	d := &DataBase{}

// 	d.Connect("postgres", dbInfo)
// 	d.Close()
// }
