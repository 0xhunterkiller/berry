package action

import "github.com/jmoiron/sqlx"

type actionStore struct {
	db *sqlx.DB
}

func NewActionStore(db *sqlx.DB) ActionStoreIface {
	return &actionStore{db: db}
}

type ActionStoreIface interface{}

var _ ActionStoreIface = &actionStore{}
