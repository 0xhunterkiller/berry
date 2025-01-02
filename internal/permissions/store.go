package permission

import (
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type permissionStore struct {
	db *sqlx.DB
}

func NewPermissionStore(db *sqlx.DB) PermissionStoreIface {
	return &permissionStore{db: db}
}

func (store *permissionStore) createPermission(permission *models.PermissionModel) error {

	query := `
		INSERT INTO permissions (name, description)
		VALUES ($1, $2)
		RETURNING id, createdat
	`
	err := store.db.QueryRowx(query, permission.Name, permission.Description).Scan(&permission.ID, &permission.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *permissionStore) deletePermission(id string) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := store.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

type PermissionStoreIface interface {
	createPermission(permission *models.PermissionModel) error
	deletePermission(id string) error
}

var _ PermissionStoreIface = &permissionStore{}
