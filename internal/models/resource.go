package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ResourceModel struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdat" db:"createdat"`
}

func (resource *ResourceModel) Validate() error {
	v := validator.New()
	err := v.Struct(resource)
	if err != nil {
		return err
	}
	return nil
}
