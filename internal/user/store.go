package user

import (
	"database/sql"
	"fmt"

	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	DB *sqlx.DB
}

type UserStoreIface interface {
	CreateUser(user *models.UserModel) error
	GetByID(userid string) (*models.UserModel, error)
	GetByUsername(username string) (*models.UserModel, error)
	UpdateByID(userid string, col string, val string) error
	ActivationToggleByID(userid string, op bool) error
	DeleteByID(userid string) error
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

func (store *UserStore) CreateUser(user *models.UserModel) error {
	query := `
		INSERT INTO users (username, email, hpassword, isactive)
		VALUES ($1, $2, $3, $4)
		RETURNING userid, createdat, updatedat
	`
	err := store.DB.QueryRowx(query, user.Username, user.Email, user.Password, user.IsActive).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (store *UserStore) GetByID(userid string) (*models.UserModel, error) {
	query := `
		SELECT userid, username, email, hpassword, isactive, createdat, updatedat
		FROM users WHERE userid = $1
	`
	var user models.UserModel
	err := store.DB.Get(&user, query, userid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (store *UserStore) GetByUsername(username string) (*models.UserModel, error) {
	query := `
		SELECT userid, username, email, hpassword, isactive, createdat, updatedat
		FROM users WHERE username = $1
	`
	var user models.UserModel
	err := store.DB.Get(&user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (store *UserStore) UpdateByID(userid string, col string, val string) error {
	columnQueries := map[string]string{
		"email":     `UPDATE users SET email = $1, updatedat = NOW() WHERE userid = $2`,
		"hpassword": `UPDATE users SET hpassword = $1, updatedat = NOW() WHERE userid = $2`,
	}

	query, exists := columnQueries[col]
	if !exists {
		return fmt.Errorf("invalid column specified for update: %s", col)
	}

	res, err := store.DB.Exec(query, val, userid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *UserStore) ActivationToggleByID(userid string, op bool) error {
	query := `UPDATE users SET isactive = $1, updatedat = NOW() WHERE userid = $2`

	res, err := store.DB.Exec(query, op, userid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *UserStore) DeleteByID(userid string) error {
	query := `DELETE FROM users WHERE userid = $1`
	_, err := store.DB.Exec(query, userid)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
