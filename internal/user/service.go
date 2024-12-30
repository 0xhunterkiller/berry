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

func (us *userService) createUser(username string, email string, password string, isactive bool) error {
	hpassword, err := validateAndGeneratePasswordHash(password)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	// setup and validate the users
	user := models.UserModel{
		Username: username,
		Email:    email,
		Password: hpassword,
		IsActive: isactive,
	}
	err = user.Validate()
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	// create the user in db
	err = us.store.createUser(&user)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *userService) getByUsername(username string) (*models.UserModel, error) {
	user, err := us.store.getByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) getByID(userID string) (*models.UserModel, error) {
	user, err := us.store.getByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) updateEmail(userID string, email string) error {

	v := validator.New()
	err := v.Var(email, "email")
	if err != nil {
		return fmt.Errorf("error validating email: %w", err)
	}

	err = us.store.updateByID(userID, "email", email)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *userService) updatePassword(userID string, password string) error {

	hpassword, err := validateAndGeneratePasswordHash(password)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = us.store.updateByID(userID, "hpassword", hpassword)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *userService) deactivateUser(userID string) error {
	err := us.store.activationToggleByID(userID, false)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *userService) activateUser(userID string) error {
	err := us.store.activationToggleByID(userID, true)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *userService) deleteUser(userID string) error {
	err := us.store.deleteByID(userID)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

type UserServiceIface interface {
	createUser(username string, email string, password string, isactive bool) error
	getByUsername(username string) (*models.UserModel, error)
	getByID(userID string) (*models.UserModel, error)
	updateEmail(userID string, email string) error
	updatePassword(userID string, password string) error
	deactivateUser(userID string) error
	activateUser(userID string) error
	deleteUser(userID string) error
}

var _ UserServiceIface = &userService{}
