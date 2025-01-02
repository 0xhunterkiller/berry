package action

import (
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type actionStore struct {
	db *sqlx.DB
}

func NewActionStore(db *sqlx.DB) ActionStoreIface {
	return &actionStore{db: db}
}

func (store *actionStore) createAction(action *models.ActionModel) error {

	query := `
		INSERT INTO actions (name, description)
		VALUES ($1, $2)
		RETURNING id, createdat
	`
	err := store.db.QueryRowx(query, action.Name, action.Description).Scan(&action.ID, &action.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *actionStore) deleteAction(id string) error {
	query := `DELETE FROM actions WHERE id = $1`
	_, err := store.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

type ActionStoreIface interface {
	createAction(action *models.ActionModel) error
	deleteAction(id string) error
}

var _ ActionStoreIface = &actionStore{}
