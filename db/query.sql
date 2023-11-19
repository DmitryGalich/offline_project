-- name: GetUsers :many
SELECT * FROM "users";

-- name: GetUser :one
SELECT * FROM "users" 
WHERE "id" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "users" (
  "login"
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateUser :one
UPDATE "users"
SET
login = COALESCE($2, login)
WHERE "id" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1;
