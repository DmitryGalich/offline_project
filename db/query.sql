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
  "date_of_birth",
  "created_at"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateUser :one
UPDATE "users"
SET
login = COALESCE($2, login)
password = COALESCE($3, password)
mail = COALESCE($4, mail)
phone = COALESCE($5, phone)
first_name = COALESCE($6, first_name)
second_name = COALESCE($7, second_name)
date_of_birth = COALESCE($7, date_of_birth)
created_at = COALESCE($7, created_at)
WHERE "id" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1;
