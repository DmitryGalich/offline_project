package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	//var err error

	postgres_name := os.Getenv("POSTGRES_USER")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_port := os.Getenv("POSTGRES_PORT")

	result := "host=localhost user=" + postgres_user + " password=" + postgres_password + " dbname=" + postgres_name + " port=" + postgres_port + " sslmode=disable"

	DB, err := gorm.Open(postgres.Open(result), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	log.Printf("KEK %s", DB.Name())
}
