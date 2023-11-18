// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createChat = `-- name: CreateChat :one
INSERT INTO "chats" DEFAULT VALUES RETURNING id
`

func (q *Queries) CreateChat(ctx context.Context) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createChat)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (
  "login"
) VALUES (
  $1
)
RETURNING id, login
`

func (q *Queries) CreateUser(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, login)
	var i User
	err := row.Scan(&i.ID, &i.Login)
	return i, err
}

const deleteChat = `-- name: DeleteChat :exec
DELETE FROM "chats"
WHERE "id" = $1
`

func (q *Queries) DeleteChat(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteChat, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getChat = `-- name: GetChat :one
SELECT id FROM "chats"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetChat(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getChat, id)
	err := row.Scan(&id)
	return id, err
}

const getChats = `-- name: GetChats :many

SELECT id FROM "chats"
`

// ----------------------
func (q *Queries) GetChats(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getChats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, login FROM "users"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Login)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, login FROM "users"
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Login); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE "users"
  set "login" = $2
WHERE id = $1
`

type UpdateUserParams struct {
	ID    uuid.UUID
	Login string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.ID, arg.Login)
	return err
}
