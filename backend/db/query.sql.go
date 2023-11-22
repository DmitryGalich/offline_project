// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
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
RETURNING id, login, password, mail, phone, first_name, second_name, date_of_birth, created_at
`

type CreateUserParams struct {
	Login       string
	Password    string
	Mail        string
	Phone       string
	FirstName   string
	SecondName  string
	DateOfBirth time.Time
	CreatedAt   time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Login,
		arg.Password,
		arg.Mail,
		arg.Phone,
		arg.FirstName,
		arg.SecondName,
		arg.DateOfBirth,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Password,
		&i.Mail,
		&i.Phone,
		&i.FirstName,
		&i.SecondName,
		&i.DateOfBirth,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "id" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, login, password, mail, phone, first_name, second_name, date_of_birth, created_at FROM "users" 
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Password,
		&i.Mail,
		&i.Phone,
		&i.FirstName,
		&i.SecondName,
		&i.DateOfBirth,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, login, password, mail, phone, first_name, second_name, date_of_birth, created_at FROM "users"
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
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Password,
			&i.Mail,
			&i.Phone,
			&i.FirstName,
			&i.SecondName,
			&i.DateOfBirth,
			&i.CreatedAt,
		); err != nil {
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

const updateUser = `-- name: UpdateUser :one
UPDATE "users"
SET
login = COALESCE($2, login),
password = COALESCE($3, password),
mail = COALESCE($4, mail),
phone = COALESCE($5, phone),
first_name = COALESCE($6, first_name),
second_name = COALESCE($7, second_name),
date_of_birth = COALESCE($7, date_of_birth),
created_at = COALESCE($8, created_at)
WHERE "id" = $1
RETURNING id, login, password, mail, phone, first_name, second_name, date_of_birth, created_at
`

type UpdateUserParams struct {
	ID         uuid.UUID
	Login      string
	Password   string
	Mail       string
	Phone      string
	FirstName  string
	SecondName string
	CreatedAt  time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Login,
		arg.Password,
		arg.Mail,
		arg.Phone,
		arg.FirstName,
		arg.SecondName,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Password,
		&i.Mail,
		&i.Phone,
		&i.FirstName,
		&i.SecondName,
		&i.DateOfBirth,
		&i.CreatedAt,
	)
	return i, err
}
