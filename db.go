package praypi

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func dbConnect(user string, pass string, dbname string, host string, port string) *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, pass, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func dbInit(db *sql.DB) error {
	stmt := `
	CREATE TABLE prayers(
		id SERIAL,
		type text,
		lang text,
		uuid text)`

	_, err := db.Exec(stmt)

	return err
}
