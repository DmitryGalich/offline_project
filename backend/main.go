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

	log.Println("CREATING USER")
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

	log.Println("UPDATING USER")
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

		local_user, err := offline_project_db.UpdateUser(context.Background(), params)
		if err != nil {
			log.Fatal("Cannot UpdateUser")
		}

		log.Println(local_user)
	}

	log.Println("GETTING USER")
	{
		local_user, err := offline_project_db.GetUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot GetUser")
		}

		log.Println("USERS: ", local_user)
	}

	log.Println("CREATING CHAT")
	chat, err := offline_project_db.CreateChat(context.Background(), "kek_chat")
	if err != nil {
		log.Fatal("Cannot CreateChat")
	}
	log.Println("CHAT: ", chat)

	log.Println("UPDATING CHAT")
	{
		params := db.UpdateChatParams{
			ID:    chat.ID,
			Title: "KEK_chat",
		}
		local_chat, err := offline_project_db.UpdateChat(context.Background(), params)
		if err != nil {
			log.Fatal("Cannot UpdateChat")
		}
		log.Println("CHAT: ", local_chat)
	}

	log.Println("GETTING CHAT")
	{
		local_chat, err := offline_project_db.GetChat(context.Background(), chat.ID)
		if err != nil {
			log.Fatal("Cannot GetChat")
		}

		log.Println("CHAT: ", local_chat)
	}

	log.Println("CREATING CHAT COMMENT")
	chat_comment_param := db.CreateChatCommentParams{
		ChatID:   chat.ID,
		AuthorID: user.ID,
		Text:     "kek",
	}
	chat_comment, err := offline_project_db.CreateChatComment(context.Background(), chat_comment_param)
	if err != nil {
		log.Fatal("Cannot CreateChatComment")
	}
	log.Println("CHAT COMMENT: ", chat_comment)

	log.Println("GETTING CHAT COMMENT")
	{
		local_chat_comment, err := offline_project_db.GetChatComment(context.Background(), chat_comment.ID)
		if err != nil {
			log.Fatal("Cannot GetChatComment")
		}

		log.Println("CHAT_COMMENT: ", local_chat_comment)
	}

	log.Println("DELETING CHAT COMMENTS")
	{
		err := offline_project_db.DeleteChatComment(context.Background(), chat_comment.ID)
		if err != nil {
			log.Fatal("Cannot DeleteChatComment")
		}
	}

	log.Println("GETTING CHAT COMMENTS")
	{
		local_chat_comments, err := offline_project_db.GetChatComments(context.Background())
		if err != nil {
			log.Fatal("Cannot GetChatComments")
		}

		log.Println("CHAT_COMMENTS: ", local_chat_comments)
	}

	log.Println("DELETING CHAT")
	{
		err := offline_project_db.DeleteChat(context.Background(), chat.ID)
		if err != nil {
			log.Fatal("Cannot DeleteChat")
		}
	}

	log.Println("GETTING CHATS")
	{
		local_chats, err := offline_project_db.GetChats(context.Background())
		if err != nil {
			log.Fatal("Cannot GetChats")
		}

		log.Println("CHATS: ", local_chats)
	}

	log.Println("DELETING USERS")
	{
		err := offline_project_db.DeleteUser(context.Background(), user.ID)
		if err != nil {
			log.Fatal("Cannot DeleteUser")
		}
	}

	log.Println("GETTING USERS")
	{
		users, err := offline_project_db.GetUsers(context.Background())
		if err != nil {
			log.Fatal("Cannot GetUsers")
		}

		log.Println("USERS: ", users)
	}
}
