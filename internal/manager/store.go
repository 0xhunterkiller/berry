package manager

import "github.com/jmoiron/sqlx"

type managerStore struct {
	db *sqlx.DB
}

func NewManagerStore(db *sqlx.DB) ManagerStoreIface {
	return &managerStore{db: db}
}

func (store *managerStore) createUserRole(user string, role string) (string, error) {
	return "", nil
}

func (store *managerStore) deleteUserRole(userRole string) error {
	return nil
}

func (store *managerStore) createRolePermission(role string, permission string) (string, error) {
	return "", nil
}

func (store *managerStore) deleteRolePermission(rolePermission string) error {
	return nil
}

func (store *managerStore) createResourceAction(resource string, action string) (string, error) {
	return "", nil
}

func (store *managerStore) deleteResourceAction(resourceAction string) error {
	return nil
}

func (store *managerStore) createPermissionResourceAction(permission string, resourceAction string) (string, error) {
	return "", nil
}

func (store *managerStore) deletePermissionResourceAction(permissionResourceAction string) error {
	return nil
}

type ManagerStoreIface interface {
	createUserRole(user string, role string) (string, error)
	deleteUserRole(userRole string) error

	createRolePermission(role string, permission string) (string, error)
	deleteRolePermission(rolePermission string) error

	createResourceAction(resource string, action string) (string, error)
	deleteResourceAction(resourceAction string) error

	createPermissionResourceAction(permission string, resourceAction string) (string, error)
	deletePermissionResourceAction(permissionResourceAction string) error
}

var _ ManagerStoreIface = &managerStore{}
