package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbname := "dbname=" + os.Getenv("POSTGRES_DB_NAME")
	username := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	host := "host=" + os.Getenv("DOCKER_DB_CONTAINER_NAME") + "." + os.Getenv("DOCKER_NETWORK_NAME")
	port := "port=" + os.Getenv("POSTGRES_PORT")
	sslmode := "sslmode=disable"

	connection_data := dbname + " " + username + " " + password + " " + host + " " + port + " " + sslmode

	fmt.Println(connection_data)

	db, err := sql.Open("postgres", connection_data)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
