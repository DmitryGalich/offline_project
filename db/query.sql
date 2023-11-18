-- name: GetUsers :many
SELECT * FROM "users";

-- name: GetUser :one
SELECT * FROM "users"
WHERE "id" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "users" (
  "id", "login"
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "users"
  set "login" = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1;

-- name: GetChats :many
SELECT * FROM "chats";

-- name: GetChat :one
SELECT * FROM "chats"
WHERE "id" = $1 LIMIT 1;

-- name: CreateChat :one
INSERT INTO "chats" DEFAULT VALUES RETURNING *;

-- name: DeleteChat :exec
DELETE FROM "chats"
WHERE "id" = $1;