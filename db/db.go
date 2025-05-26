package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	db *sql.DB
}

func (db Database) GetConnection() *sql.DB {

	if db.db == nil {
		fmt.Println(os.Getenv("DBUSER"), os.Getenv("DBPASS"))
		connectionString := fmt.Sprintf("postgres://%s:%s@postgresql:5432/postgres", os.Getenv("DBUSER"), os.Getenv("DBPASS"))
		conn, err := sql.Open("pgx", connectionString)
		if err != nil {
			fmt.Println(err)
		}
		db.db = conn
	}
	return db.db
}
