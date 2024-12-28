package appinit

import (
	"github.com/0xhunterkiller/berry/internal/auth"
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/internal/user"
)

type Services struct {
	UserService *user.UserService
}

func initializeServices(inj *models.Deps) *Services {
	userStore := user.NewUserStore(inj.DB)

	return &Services{
		UserService: user.NewUserService(userStore),
	}
}

type Handlers struct {
	AuthHandler *auth.AuthHandler
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		AuthHandler: auth.NewAuthHandler(services.UserService),
	}
}

func AppInit(inj *models.Deps) *Handlers {
	services := initializeServices(inj)
	handlers := initializeHandlers(services)
	return handlers
}
