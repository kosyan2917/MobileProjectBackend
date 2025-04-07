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
		connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5431/postgres", os.Getenv("DBUSER"), os.Getenv("DBPASS"))
		print(connectionString)
		conn, err := sql.Open("pgx", connectionString)
		if err != nil {
			fmt.Println(err)
		}
		db.db = conn
	}
	return db.db
}
