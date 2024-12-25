package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserModel struct {
	// System Generated
	CreatedAt time.Time `json:"createdat" db:"createdat"`
	UpdatedAt time.Time `json:"updatedat" db:"updatedat"`
	ID        string    `json:"userid" db:"userid"`

	// User Input
	Username string `json:"username" db:"username" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"hpassword" db:"hpassword" validate:"required"`
	IsActive bool   `json:"isactive" db:"isactive" validate:"required"`
}

func (user *UserModel) Validate() error {

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		return fmt.Errorf("an error occured while validating the user: %w", err)
	}

	return nil
}
