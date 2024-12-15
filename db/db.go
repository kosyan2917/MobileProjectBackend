package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

type Database struct {
	db *sql.DB
}

func (db Database) GetConnection() *sql.DB {

	if db.db == nil {
		connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres", os.Getenv("DBUSER"), os.Getenv("DBPASS"))
		print(connectionString)
		conn, err := sql.Open("pgx", connectionString)
		if err != nil {
			fmt.Println(err)
		}
		db.db = conn
	}
	return db.db
}
