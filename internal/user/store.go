package user

import (
	"database/sql"
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type userStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) UserStoreIface {
	return &userStore{db: db}
}

func (store *userStore) createUser(user *models.UserModel) error {
	query := `
		INSERT INTO users (username, email, hpassword, isactive)
		VALUES ($1, $2, $3, $4)
		RETURNING userid, createdat, updatedat
	`
	err := store.db.QueryRowx(query, user.Username, user.Email, user.Password, user.IsActive).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (store *userStore) getByID(userid string) (*models.UserModel, error) {
	query := `
		SELECT userid, username, email, hpassword, isactive, createdat, updatedat
		FROM users WHERE userid = $1
	`
	var user models.UserModel
	err := store.db.Get(&user, query, userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (store *userStore) getByUsername(username string) (*models.UserModel, error) {
	query := `
		SELECT userid, username, email, hpassword, isactive, createdat, updatedat
		FROM users WHERE username = $1
	`
	var user models.UserModel
	err := store.db.Get(&user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (store *userStore) updateByID(userid string, col string, val string) error {
	columnQueries := map[string]string{
		"email":     `UPDATE users SET email = $1, updatedat = NOW() WHERE userid = $2`,
		"hpassword": `UPDATE users SET hpassword = $1, updatedat = NOW() WHERE userid = $2`,
	}

	query, exists := columnQueries[col]
	if !exists {
		return fmt.Errorf("invalid column specified for update: %s", col)
	}

	res, err := store.db.Exec(query, val, userid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *userStore) activationToggleByID(userid string, op bool) error {
	query := `UPDATE users SET isactive = $1, updatedat = NOW() WHERE userid = $2`

	res, err := store.db.Exec(query, op, userid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *userStore) deleteByID(userid string) error {
	query := `DELETE FROM users WHERE userid = $1`
	_, err := store.db.Exec(query, userid)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

type UserStoreIface interface {
	createUser(user *models.UserModel) error
	getByID(userid string) (*models.UserModel, error)
	getByUsername(username string) (*models.UserModel, error)
	updateByID(userid string, col string, val string) error
	activationToggleByID(userid string, op bool) error
	deleteByID(userid string) error
}

var _ UserStoreIface = &userStore{}
