package auth

import (
	"context"
	"log/slog"
	"offline_project/internal/domain/models"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(
		ctx context.Context,
		email string,
	) (uid int64, err error)

	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)
}

type AppProvider interface {
	App(
		ctx context.Context,
		appID int,
	) (models.App, error)
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}
