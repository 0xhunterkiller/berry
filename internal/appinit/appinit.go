package appinit

import (
	"github.com/0xhunterkiller/berry/internal/auth"
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/internal/user"
)

type Services struct {
	UserService *user.UserService
}

func initializeServices(ston *models.Deps) *Services {
	userStore := user.NewUserStore(ston.DB)

	return &Services{
		UserService: user.NewUserService(userStore),
	}
}

type Handlers struct {
	AuthHandler *auth.AuthHandler
}

func initializeHandlers(services *Services, ston *models.Deps) *Handlers {
	return &Handlers{
		AuthHandler: auth.NewAuthHandler(services.UserService, ston.Logger),
	}
}

func AppInit(ston *models.Deps) *Handlers {
	services := initializeServices(ston)
	handlers := initializeHandlers(services, ston)
	return handlers
}
