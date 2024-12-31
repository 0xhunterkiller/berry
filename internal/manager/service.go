package manager

type managerService struct {
	store ManagerStoreIface
}

func NewManagerService(store ManagerStoreIface) ManagerServiceIface {
	return &managerService{store: store}
}

func (ms *managerService) createUserRole(user string, role string) (string, error) {
	return "", nil
}

func (ms *managerService) deleteUserRole(userRole string) error {
	return nil
}

func (ms *managerService) createRolePermission(role string, permission string) (string, error) {
	return "", nil
}

func (ms *managerService) deleteRolePermission(rolePermission string) error {
	return nil
}

func (ms *managerService) createInteraction(resource string, action string) (string, error) {
	return "", nil
}

func (ms *managerService) deleteInteraction(interaction string) error {
	return nil
}

func (ms *managerService) createPermissionInteraction(permission string, interaction string) (string, error) {
	return "", nil
}

func (ms *managerService) deletePermissionInteraction(permissionInteraction string) error {
	return nil
}

type ManagerServiceIface interface {
	// All Params are IDs (UUID)
	createUserRole(user string, role string) (string, error)
	deleteUserRole(userRole string) error

	createRolePermission(role string, permission string) (string, error)
	deleteRolePermission(rolePermission string) error

	createInteraction(resource string, action string) (string, error)
	deleteInteraction(interaction string) error

	createPermissionInteraction(permission string, interaction string) (string, error)
	deletePermissionInteraction(permissionInteraction string) error
}

var _ ManagerServiceIface = &managerService{}
