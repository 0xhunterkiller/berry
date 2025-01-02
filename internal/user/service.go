package user

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/go-playground/validator/v10"
	pv "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	store UserStoreIface
}

func NewUserService(store UserStoreIface) UserServiceIface {
	return &userService{store: store}
}

func validateAndGeneratePasswordHash(password string) (string, error) {
	// validatePassword checks for password complexity.
	const minEntropyBits = 60
	err := pv.Validate(password, minEntropyBits)
	if err != nil {
		return "", fmt.Errorf("password too weak: %v", err.Error())
	}

	// hash the password with bcrypt algorithm, which is suitable for passwords at rest
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("an error occured while hashing the password: %w", err)
	}
	return string(hashedPassword), nil
}

func (us *userService) createUser(username string, email string, password string, isactive bool) (string, error) {
	password, err := validateAndGeneratePasswordHash(password)
	if err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}
	// setup and validate the users
	user := models.UserModel{
		Username: username,
		Email:    email,
		Password: password,
		IsActive: isactive,
	}
	err = user.Validate()
	if err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}
	// create the user in db
	err = us.store.createUser(&user)
	if err != nil {
		return "", fmt.Errorf("error while committing user to db: %w", err)
	}

	if user.ID == "" {
		return "", fmt.Errorf("error while committing user to db: %w", err)
	}
	return user.ID, nil
}

func (us *userService) getByUsername(username string) (*models.UserModel, error) {
	user, err := us.store.getByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) getByID(id string) (*models.UserModel, error) {
	user, err := us.store.getByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) updateEmail(id string, email string) error {

	v := validator.New()
	err := v.Var(email, "email")
	if err != nil {
		return fmt.Errorf("error validating email: %w", err)
	}

	err = us.store.updateByID(id, "email", email)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *userService) updatePassword(id string, password string) error {

	password, err := validateAndGeneratePasswordHash(password)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = us.store.updateByID(id, "password", password)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *userService) deactivateUser(id string) error {
	err := us.store.activationToggleByID(id, false)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *userService) activateUser(id string) error {
	err := us.store.activationToggleByID(id, true)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *userService) deleteUser(id string) error {
	err := us.store.deleteByID(id)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

type UserServiceIface interface {
	createUser(username string, email string, password string, isactive bool) (string, error)
	getByUsername(username string) (*models.UserModel, error)
	getByID(id string) (*models.UserModel, error)
	updateEmail(id string, email string) error
	updatePassword(id string, password string) error
	deactivateUser(id string) error
	activateUser(id string) error
	deleteUser(id string) error
}

var _ UserServiceIface = &userService{}
