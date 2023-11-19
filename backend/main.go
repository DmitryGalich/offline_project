package main

import (
	"context"
	"database/sql"
	"log"
	"offline_project/db"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db_connecting_data := os.Getenv("POSTGRES_CONNECTION_STRING")
	if db_connecting_data == "" {
		log.Fatal("db_connecting_data is empty")
	}

	conn, err := sql.Open("postgres", db_connecting_data)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer conn.Close()

	if conn.Ping() != nil {
		log.Fatal("Cannot ping database")
	}

	offline_project_db := db.New(conn)

	user, err := offline_project_db.CreateUser(context.Background(), "LEL")
	if err != nil {
		log.Fatal("Cannot CreateUser")
	}

	log.Println(user)
}
