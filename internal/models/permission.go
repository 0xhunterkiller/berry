package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type PermissionModel struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdat" db:"createdat"`
}

func (permission *PermissionModel) Validate() error {
	v := validator.New()
	err := v.Struct(permission)
	if err != nil {
		return err
	}
	return nil
}
