package manager

import (
	"github.com/jmoiron/sqlx"
)

type managerStore struct {
	db *sqlx.DB
}

func NewManagerStore(db *sqlx.DB) ManagerStoreIface {
	return &managerStore{db: db}
}

func (store *managerStore) createUserRole(user string, role string) (string, error) {
	var id string
	query := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) RETURNING id`
	err := store.db.QueryRowx(query, user, role).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (store *managerStore) deleteUserRole(userRole string) error {
	query := `DELETE FROM users_roles WHERE id = $1`
	_, err := store.db.Exec(query, userRole)
	if err != nil {
		return err
	}
	return nil
}

func (store *managerStore) createRolePermission(role string, permission string) (string, error) {
	var id string
	query := `INSERT INTO roles_permissions (role_id, permission_id) VALUES ($1, $2) RETURNING id`
	err := store.db.QueryRowx(query, role, permission).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (store *managerStore) deleteRolePermission(rolePermission string) error {
	query := `DELETE FROM roles_permissions WHERE id = $1`
	_, err := store.db.Exec(query, rolePermission)
	if err != nil {
		return err
	}
	return nil
}

func (store *managerStore) createInteraction(resource string, action string) (string, error) {
	var id string
	query := `INSERT INTO interactions (resource_id, action_id) VALUES ($1, $2) RETURNING id`
	err := store.db.QueryRowx(query, resource, action).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (store *managerStore) deleteInteraction(interaction string) error {
	query := `DELETE FROM interactions WHERE id = $1`
	_, err := store.db.Exec(query, interaction)
	if err != nil {
		return err
	}
	return nil
}

func (store *managerStore) createPermissionInteraction(permission string, interaction string) (string, error) {
	var id string
	query := `INSERT INTO permissions_interactions (permission_id, interaction_id) VALUES ($1, $2) RETURNING id`
	err := store.db.QueryRowx(query, permission, interaction).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (store *managerStore) deletePermissionInteraction(permissionInteraction string) error {
	query := `DELETE FROM permissions_interactions WHERE id = $1`
	_, err := store.db.Exec(query, permissionInteraction)
	if err != nil {
		return err
	}
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
