package permissions

import "github.com/jmoiron/sqlx"

type permissionStore struct {
	db *sqlx.DB
}

func NewPermissionStore(db *sqlx.DB) PermissionStoreIface {
	return &permissionStore{db: db}
}

type PermissionStoreIface interface{}

var _ PermissionStoreIface = &permissionStore{}
