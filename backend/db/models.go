// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/google/uuid"
)

type Chat struct {
	ID uuid.UUID
}

type User struct {
	ID    uuid.UUID
	Login string
}
