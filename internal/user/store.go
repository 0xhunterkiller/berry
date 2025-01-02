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
		INSERT INTO users (name, email, password, isactive)
		VALUES ($1, $2, $3, $4)
		RETURNING id, createdat, updatedat
	`
	err := store.db.QueryRowx(query, user.Username, user.Email, user.Password, user.IsActive).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *userStore) getByID(id string) (*models.UserModel, error) {
	query := `
		SELECT id, name, email, password, isactive, createdat, updatedat
		FROM users WHERE id = $1
	`
	var user models.UserModel
	err := store.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (store *userStore) getByUsername(name string) (*models.UserModel, error) {
	query := `
		SELECT id, name, email, password, isactive, createdat, updatedat
		FROM users WHERE name = $1
	`
	var user models.UserModel
	err := store.db.Get(&user, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (store *userStore) updateByID(id string, col string, val string) error {
	columnQueries := map[string]string{
		"email":    `UPDATE users SET email = $1, updatedat = NOW() WHERE id = $2`,
		"password": `UPDATE users SET password = $1, updatedat = NOW() WHERE id = $2`,
	}

	query, exists := columnQueries[col]
	if !exists {
		return fmt.Errorf("invalid column specified for update: %s", col)
	}

	res, err := store.db.Exec(query, val, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *userStore) activationToggleByID(id string, op bool) error {
	query := `UPDATE users SET isactive = $1, updatedat = NOW() WHERE id = $2`

	res, err := store.db.Exec(query, op, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (store *userStore) deleteByID(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := store.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

type UserStoreIface interface {
	createUser(user *models.UserModel) error
	getByID(id string) (*models.UserModel, error)
	getByUsername(name string) (*models.UserModel, error)
	updateByID(id string, col string, val string) error
	activationToggleByID(id string, op bool) error
	deleteByID(id string) error
}

var _ UserStoreIface = &userStore{}
