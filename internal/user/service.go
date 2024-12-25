package user

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/internal/store"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, email string, password string, isactive bool) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("an error occured while hashing the password: %w", err)
	}

	user := models.UserModel{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		IsActive: isactive,
	}

	err = user.Validate()
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = store.Store.UserStore.CreateUser(&user)
	if err != nil {
		return fmt.Errorf("error while commiting new user to db: %w", err)
	}

	return nil
}
