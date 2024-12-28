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
	UpdateByID(user *models.UserModel) error
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

func (store *UserStore) UpdateByID(user *models.UserModel) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, hpassword = $3, isactive = $4, updatedat = NOW()
		WHERE userid = $5
	`
	res, err := store.DB.Exec(query, user.Username, user.Email, user.Password, user.IsActive, user.ID)
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
