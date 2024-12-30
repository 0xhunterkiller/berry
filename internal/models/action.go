package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ActionModel struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdat" db:"createdat"`
}

func (action *ActionModel) Validate() error {
	v := validator.New()
	err := v.Struct(action)
	if err != nil {
		return err
	}
	return nil
}
