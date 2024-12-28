package user

import (
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
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
}

func NewUserService(userStore *UserStore) *UserService {
	return &UserService{Store: userStore}
}

func (us *UserService) CreateUser(username string, email string, password string, isactive bool) error {

	// validatePassword checks for password complexity.
	const minEntropyBits = 60
	err := pv.Validate(password, minEntropyBits)
	if err != nil {
		return fmt.Errorf("password too weak: %v", err.Error())
	}

	// hash the password with bcrypt algorithm, which is suitable for passwords at rest
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("an error occured while hashing the password: %w", err)
	}

	// setup and validate the users
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
