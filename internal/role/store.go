package role

import "github.com/jmoiron/sqlx"

type roleStore struct {
	db *sqlx.DB
}

func NewRoleStore(db *sqlx.DB) RoleStoreIface {
	return &roleStore{db: db}
}

type RoleStoreIface interface{}

var _ RoleStoreIface = &roleStore{}
