package appinit

import (
	"github.com/0xhunterkiller/berry/internal/action"
	"github.com/0xhunterkiller/berry/internal/manager"
	"github.com/0xhunterkiller/berry/internal/models"
	permission "github.com/0xhunterkiller/berry/internal/permissions"
	"github.com/0xhunterkiller/berry/internal/resource"
	"github.com/0xhunterkiller/berry/internal/role"
	"github.com/0xhunterkiller/berry/internal/user"
)

type Services struct {
	UserService       user.UserServiceIface
	RoleService       role.RoleServiceIface
	ActionService     action.ActionServiceIface
	PermissionService permission.PermissionServiceIface
	ResourceService   resource.ResourceServiceIface
	ManagerService    manager.ManagerServiceIface
}

func initializeServices(inj *models.Deps) *Services {
	userStore := user.NewUserStore(inj.DB)
	roleStore := role.NewRoleStore(inj.DB)
	actionStore := action.NewActionStore(inj.DB)
	permissionStore := permission.NewPermissionStore(inj.DB)
	resourceStore := resource.NewResourceStore(inj.DB)
	managerStore := manager.NewManagerStore(inj.DB)

	return &Services{
		UserService:       user.NewUserService(userStore),
		RoleService:       role.NewRoleService(roleStore),
		ActionService:     action.NewActionService(actionStore),
		PermissionService: permission.NewPermissionService(permissionStore),
		ResourceService:   resource.NewResourceService(resourceStore),
		ManagerService:    manager.NewManagerService(managerStore),
	}
}

type Handlers struct {
	UserHandler       user.UserHandlerIface
	RoleHandler       role.RoleHandlerIface
	ActionHandler     action.ActionHandlerIface
	PermissionHandler permission.PermissionHandlerIface
	ResourceHandler   resource.ResourceHandlerIface
	ManagerHandler    manager.ManagerHandlerIface
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler:       user.NewUserHandler(services.UserService),
		RoleHandler:       role.NewRoleHandler(services.RoleService),
		ActionHandler:     action.NewActionHandler(services.ActionService),
		PermissionHandler: permission.NewPermissionHandler(services.PermissionService),
		ResourceHandler:   resource.NewResourceHandler(services.ResourceService),
		ManagerHandler:    manager.NewManagerHandler(services.ManagerService),
	}
}

func AppInit(inj *models.Deps) *Handlers {
	services := initializeServices(inj)
	handlers := initializeHandlers(services)
	return handlers
}
