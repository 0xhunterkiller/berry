package appinit

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	pv "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdminUser(db *sqlx.DB, email string, password string) (string, error) {
	query := `INSERT INTO users (username, email, hpassword, isactive) VALUES ($1, $2, $3, $4) RETURNING userid`

	// validatePassword checks for password complexity.
	const minEntropyBits = 120
	err := pv.Validate(password, minEntropyBits)
	if err != nil {
		return "", fmt.Errorf("admin password too weak: %v", err.Error())
	}

	// hash the password with bcrypt algorithm, which is suitable for passwords at rest
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("an error occured while hashing the password: %w", err)
	}

	var userid string
	err = db.QueryRowx(query, "berryroot", email, string(hashedPassword), true).Scan(&userid)
	if err != nil {
		return "", fmt.Errorf("failed to create admin user: %w", err)
	}
	return userid, nil
}

func CreateRootRole(db *sqlx.DB) (string, error) {
	query := `INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING id`

	var id string
	err := db.QueryRowx(query, "root", "this soul will have all access there is").Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func MakeRoot(db *sqlx.DB, adminUserID string, adminRoleID string) error {
	query := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) RETURNING id`

	var id string
	err := db.QueryRowx(query, adminUserID, adminRoleID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
