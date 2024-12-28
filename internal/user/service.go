package user

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/go-playground/validator/v10"
	pv "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Store UserStoreIface
}

type UserServiceIface interface {
	CreateUser(username string, email string, password string, isactive bool) error
	GetByUsername(username string) (*models.UserModel, error)
	GetByID(userID string) (*models.UserModel, error)
	UpdateEmail(userID string, email string) error
	UpdatePassword(userID string, password string) error
	DeactivateUser(userID string) error
	ActivateUser(userID string) error
	DeleteUser(userID string) error
}

func NewUserService(userStore *UserStore) *UserService {
	return &UserService{Store: userStore}
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

func (us *UserService) CreateUser(username string, email string, password string, isactive bool) error {

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
	err = us.Store.CreateUser(&user)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *UserService) GetByUsername(username string) (*models.UserModel, error) {
	user, err := us.Store.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetByID(userID string) (*models.UserModel, error) {
	user, err := us.Store.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateEmail(userID string, email string) error {

	v := validator.New()
	err := v.Var(email, "email")
	if err != nil {
		return fmt.Errorf("error validating email: %w", err)
	}

	err = us.Store.UpdateByID(userID, "email", email)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *UserService) UpdatePassword(userID string, password string) error {

	hpassword, err := validateAndGeneratePasswordHash(password)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	err = us.Store.UpdateByID(userID, "hpassword", hpassword)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}

	return nil
}

func (us *UserService) DeactivateUser(userID string) error {
	err := us.Store.ActivationToggleByID(userID, false)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *UserService) ActivateUser(userID string) error {
	err := us.Store.ActivationToggleByID(userID, true)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}

func (us *UserService) DeleteUser(userID string) error {
	err := us.Store.DeleteByID(userID)
	if err != nil {
		return fmt.Errorf("error while committing user to db: %w", err)
	}
	return nil
}
