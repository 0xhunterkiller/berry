package appinit

import (
	"github.com/0xhunterkiller/berry/internal/auth"
	"github.com/0xhunterkiller/berry/internal/user"
	"github.com/jmoiron/sqlx"
)

type Services struct {
	UserService *user.UserService
}

func initializeServices(db *sqlx.DB) *Services {
	userStore := user.NewUserStore(db)

	return &Services{
		UserService: user.NewUserService(userStore),
	}
}

type Handlers struct {
	AuthHandler *auth.AuthHandler
}

func initializeHandlers(db *sqlx.DB) *Handlers {
	services := initializeServices(db)
	return &Handlers{
		AuthHandler: auth.NewAuthHandler(services.UserService),
	}
}

func AppInit(db *sqlx.DB) *Handlers {
	handlers := initializeHandlers(db)
	return handlers
}
