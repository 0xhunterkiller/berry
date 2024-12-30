package appinit

import (
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/internal/role"
	"github.com/0xhunterkiller/berry/internal/user"
)

type Services struct {
	UserService user.UserServiceIface
	RoleService role.RoleServiceIface
}

func initializeServices(inj *models.Deps) *Services {
	userStore := user.NewUserStore(inj.DB)
	roleStore := role.NewRoleStore(inj.DB)

	return &Services{
		UserService: user.NewUserService(userStore),
		RoleService: role.NewRoleService(roleStore),
	}
}

type Handlers struct {
	UserHandler user.UserHandlerIface
	RoleHandler role.RoleHandlerIface
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler: user.NewUserHandler(services.UserService),
		RoleHandler: role.NewRoleHandler(services.RoleService),
	}
}

func AppInit(inj *models.Deps) *Handlers {
	services := initializeServices(inj)
	handlers := initializeHandlers(services)
	return handlers
}
