package resource

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type resourceStore struct {
	db *sqlx.DB
}

func NewResourceStore(db *sqlx.DB) ResourceStoreIface {
	return &resourceStore{db: db}
}

func (store *resourceStore) createResource(resource *models.ResourceModel) error {

	query := `
		INSERT INTO resources (name, description)
		VALUES ($1, $2)
		RETURNING id, createdat
	`
	err := store.db.QueryRowx(query, resource.Name, resource.Description).Scan(&resource.ID, &resource.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (store *resourceStore) deleteResource(id string) error {
	query := `DELETE FROM resources WHERE id = $1`
	_, err := store.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

type ResourceStoreIface interface {
	createResource(resource *models.ResourceModel) error
	deleteResource(id string) error
}

var _ ResourceStoreIface = &resourceStore{}
