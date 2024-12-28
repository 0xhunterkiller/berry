package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	goaway "github.com/TwiN/go-away"
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

var reservedUsernames = []string{
	"admin", "administrator", "root", "sysadmin", "superuser", "system",
	"guest", "support", "moderator", "operator", "test", "user", "owner",
	"default", "anonymous", "unknown", "null", "temp", "example",
	"ceo", "cto", "manager", "founder", "team", "staff", "employee",
	"api", "bot", "service", "webmaster", "postmaster", "noreply", "mailer", "adminbot",
	"http", "https", "ftp", "mail", "smtp", "imap", "dns", "localhost", "berry",
}

func isReserved(username string) bool {
	username = strings.ToLower(username)
	for _, reserved := range reservedUsernames {
		if username == reserved {
			return true
		}
	}
	return false
}

func isProfane(username string) bool {
	return goaway.IsProfane(username)
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]{8,32}$`)

	if len([]byte(username)) > 240 {
		return false
	}

	return re.MatchString(username)
}

func (user *UserModel) Validate() error {

	validate := validator.New()
	username := user.Username
	if !isValidUsername(username) {
		return fmt.Errorf("cannot use this username (only alphanumeric characters, underscores, hyphens and 8-32 characters long): %v", username)
	}

	if isReserved(username) {
		return fmt.Errorf("cannot use the system reserved username: %v", username)
	}

	if isProfane(username) {
		return fmt.Errorf("cannot use this username (profanity detected): %v", username)
	}

	err := validate.Struct(user)
	if err != nil {
		return fmt.Errorf("an error occured while validating the user: %w", err)
	}

	return nil
}
