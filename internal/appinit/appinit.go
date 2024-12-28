package appinit

import (
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
	UserHandler *user.UserHandler
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler: user.NewUserHandler(services.UserService),
	}
}

func AppInit(inj *models.Deps) *Handlers {
	services := initializeServices(inj)
	handlers := initializeHandlers(services)
	return handlers
}
