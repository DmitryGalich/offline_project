package main

import (
	"context"
	"database/sql"
	"log"
	"offline_project/db"
	"os"
	"time"

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

	// Creating user

	create_user_params := db.CreateUserParams{
		Login:       "kek_login",
		Password:    "kek_password",
		Mail:        "kek_mail",
		Phone:       "kek_phone",
		FirstName:   "kek_first_name",
		SecondName:  "kek_second_name",
		DateOfBirth: time.Now(),
	}

	user, err := offline_project_db.CreateUser(context.Background(), create_user_params)
	if err != nil {
		log.Fatal("Cannot CreateUser")
	}

	log.Println(user)

	{
		params := db.UpdateUserParams{
			ID:          user.ID,
			Login:       "KEK_login",
			Password:    "KEK_password",
			Mail:        "KEK_mail",
			Phone:       "KEK_phone",
			FirstName:   "KEK_first_name",
			SecondName:  "KEK_second_name",
			DateOfBirth: time.Now(),
		}

		user, err := offline_project_db.UpdateUser(context.Background(), params)
		if err != nil {
			log.Fatal("Cannot UpdateUser")
		}

		log.Println(user)
	}

	log.Println("USER: ")

	{
		user, err := offline_project_db.GetUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot GetUser")
		}

		log.Println(user)
	}

	{
		err := offline_project_db.DeleteUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot DeleteUser")
		}
	}

	log.Println("USERS: ")

	{
		users, err := offline_project_db.GetUsers(context.Background())
		if err != nil {
			log.Fatal("Cannot GetUsers")
		}

		log.Println(users)
	}
}
