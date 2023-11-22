-- name: GetUsers :many
SELECT * FROM "users";

-- name: GetUser :one
SELECT * FROM "users" 
WHERE "id" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "users" (
  "login",
  "password",
  "mail",
  "phone",
  "first_name",
  "second_name",
  "date_of_birth"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateUser :one
UPDATE "users"
SET
login = COALESCE($2, login),
password = COALESCE($3, password),
mail = COALESCE($4, mail),
phone = COALESCE($5, phone),
first_name = COALESCE($6, first_name),
second_name = COALESCE($7, second_name),
date_of_birth = COALESCE($8, date_of_birth)
WHERE "id" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1;

-------------------------------

-- name: GetChats :many
SELECT * FROM "chats";

-- name: GetChat :one
SELECT * FROM "chats" 
WHERE "id" = $1 LIMIT 1;

-- name: CreateChat :one
INSERT INTO "chats" 
( "title" ) VALUES ( $1 )
RETURNING *;

-- name: UpdateChat :one
UPDATE "chats"
SET
title = COALESCE($2, title)
WHERE "id" = $1
RETURNING *;

-- name: DeleteChat :exec
DELETE FROM "chats"
WHERE "id" = $1;

-------------------------------

-- name: GetChatComments :many
SELECT * FROM "chat_comments";

-- name: GetChatComment :one
SELECT * FROM "chat_comments" 
WHERE "id" = $1 LIMIT 1;

-- name: CreateChatComment :one
INSERT INTO "chat_comments" (
  "chat_id",
  "author_id",
  "text"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateChatComment :one
UPDATE "chat_comments"
SET
chat_id = COALESCE($2, chat_id),
author_id = COALESCE($3, author_id),
edited_at = COALESCE($4, edited_at),
text = COALESCE($5, text)
WHERE "id" = $1
RETURNING *;

-- name: DeleteChatComment :exec
DELETE FROM "chat_comments"
WHERE "id" = $1;
