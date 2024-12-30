package role

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type roleStore struct {
	db *sqlx.DB
}

func NewRoleStore(db *sqlx.DB) RoleStoreIface {
	return &roleStore{db: db}
}

func (store *roleStore) createRole(role *models.RoleModel) error {

	query := `
		INSERT INTO roles (name, description)
		VALUES ($1, $2)
		RETURNING id, createdat
	`
	err := store.db.QueryRowx(query, role.Name, role.Description).Scan(&role.ID, &role.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (store *roleStore) deleteRole(id string) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := store.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

type RoleStoreIface interface {
	createRole(role *models.RoleModel) error
	deleteRole(id string) error
}

var _ RoleStoreIface = &roleStore{}
