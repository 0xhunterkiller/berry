package manager

type managerService struct {
	store ManagerStoreIface
}

func NewManagerService(store ManagerStoreIface) ManagerServiceIface {
	return &managerService{store: store}
}

func (ms *managerService) createUserRole(user string, role string) (string, error) {
	id, err := ms.store.createUserRole(user, role)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ms *managerService) deleteUserRole(userRole string) error {
	err := ms.store.deleteUserRole(userRole)
	if err != nil {
		return err
	}
	return nil
}

func (ms *managerService) createRolePermission(role string, permission string) (string, error) {
	id, err := ms.store.createRolePermission(role, permission)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ms *managerService) deleteRolePermission(rolePermission string) error {
	err := ms.store.deleteRolePermission(rolePermission)
	if err != nil {
		return err
	}
	return nil
}

func (ms *managerService) createInteraction(resource string, action string) (string, error) {
	id, err := ms.store.createInteraction(resource, action)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ms *managerService) deleteInteraction(interaction string) error {
	err := ms.store.deleteInteraction(interaction)
	if err != nil {
		return err
	}
	return nil
}

func (ms *managerService) createPermissionInteraction(permission string, interaction string) (string, error) {
	id, err := ms.store.createPermissionInteraction(permission, interaction)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ms *managerService) deletePermissionInteraction(permissionInteraction string) error {
	err := ms.store.deletePermissionInteraction(permissionInteraction)
	if err != nil {
		return err
	}
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
