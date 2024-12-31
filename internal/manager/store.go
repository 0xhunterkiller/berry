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

func (store *managerStore) createInteraction(resource string, action string) (string, error) {
	return "", nil
}

func (store *managerStore) deleteInteraction(interaction string) error {
	return nil
}

func (store *managerStore) createPermissionInteraction(permission string, interaction string) (string, error) {
	return "", nil
}

func (store *managerStore) deletePermissionInteraction(permissionInteraction string) error {
	return nil
}

type ManagerStoreIface interface {
	createUserRole(user string, role string) (string, error)
	deleteUserRole(userRole string) error

	createRolePermission(role string, permission string) (string, error)
	deleteRolePermission(rolePermission string) error

	createInteraction(resource string, action string) (string, error)
	deleteInteraction(interaction string) error

	createPermissionInteraction(permission string, interaction string) (string, error)
	deletePermissionInteraction(permissionInteraction string) error
}

var _ ManagerStoreIface = &managerStore{}
