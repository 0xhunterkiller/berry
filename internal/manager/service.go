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

func (ms *managerService) createResourceAction(resource string, action string) (string, error) {
	return "", nil
}

func (ms *managerService) deleteResourceAction(resourceAction string) error {
	return nil
}

func (ms *managerService) createPermissionResourceAction(permission string, resourceAction string) (string, error) {
	return "", nil
}

func (ms *managerService) deletePermissionResourceAction(permissionResourceAction string) error {
	return nil
}

type ManagerServiceIface interface {
	// All Params are IDs (UUID)
	createUserRole(user string, role string) (string, error)
	deleteUserRole(userRole string) error

	createRolePermission(role string, permission string) (string, error)
	deleteRolePermission(rolePermission string) error

	createResourceAction(resource string, action string) (string, error)
	deleteResourceAction(resourceAction string) error

	createPermissionResourceAction(permission string, resourceAction string) (string, error)
	deletePermissionResourceAction(permissionResourceAction string) error
}

var _ ManagerServiceIface = &managerService{}
