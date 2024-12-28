package models

import (
	"github.com/jmoiron/sqlx"
)

type Deps struct {
	DB *sqlx.DB
}
