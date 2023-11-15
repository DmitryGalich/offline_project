package postgresql

import (
	"database/sql"
	"os"
)

func ConnectToDB() (*sql.DB, error) {
	connection_info := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := sql.Open("postgres", connection_info)
	if err != nil {
		return nil, err
	}

	return db, nil
}
