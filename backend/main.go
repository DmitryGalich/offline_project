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
	log.Println("USER: ", user)

	// Updating user
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

	// Getting user
	{
		user, err := offline_project_db.GetUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot GetUser")
		}

		log.Println("USERS: ", user)
	}

	// Creating chat
	chat, err := offline_project_db.CreateChat(context.Background(), "kek_chat")
	if err != nil {
		log.Fatal("Cannot CreateChat")
	}
	log.Println("CHAT: ", chat)

	// Updating chat
	{
		params := db.UpdateChatParams{
			ID:    chat.ID,
			Title: "KEK_chat",
		}
		updated_chat, err := offline_project_db.UpdateChat(context.Background(), params)
		if err != nil {
			log.Fatal("Cannot CreateChat")
		}
		log.Println("CHAT: ", updated_chat)
	}

	// Getting chat
	{
		local_chat, err := offline_project_db.GetChat(context.Background(), chat.ID)
		if err != nil {
			log.Fatal("Cannot GetChat")
		}

		log.Println("CHAT: ", local_chat)
	}

	// Deleting chat
	{
		err := offline_project_db.DeleteChat(context.Background(), chat.ID)
		if err != nil {
			log.Fatal("Cannot DeleteChat")
		}
	}

	// Deleting user
	{
		err := offline_project_db.DeleteUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot DeleteUser")
		}
	}

	// Getting users
	{
		users, err := offline_project_db.GetUsers(context.Background())
		if err != nil {
			log.Fatal("Cannot GetUsers")
		}

		log.Println("USERS: ", users)
	}
}
