package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connection_data := "postgres://" + username + ":" + password + "@localhost:" +

		//const conString = "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"

		//fmt.Println(connection_data)
}
