package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connection_string := os.Getenv("POSTGRES_CONNECTION_STRING")

	fmt.Println(connection_string)

	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
