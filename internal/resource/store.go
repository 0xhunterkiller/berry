package resource

import "github.com/jmoiron/sqlx"

type resourceStore struct {
	db *sqlx.DB
}

func NewResourceStore(db *sqlx.DB) ResourceStoreIface {
	return &resourceStore{db: db}
}

type ResourceStoreIface interface{}

var _ ResourceStoreIface = &resourceStore{}
