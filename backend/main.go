package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	connection_data := "postgres://" + username + ":" + password + "@localhost:" + port + "/" + dbname

	//const conString = "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"

	fmt.Println(connection_data)
}
