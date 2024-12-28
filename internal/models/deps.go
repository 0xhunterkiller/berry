package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Deps struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}
