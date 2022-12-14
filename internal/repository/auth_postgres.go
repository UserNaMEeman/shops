package repository

import (
	"context"
	"fmt"

	"github.com/UserNaMEeman/shops/app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user app.User) (string, error) {
	var existUser string
	var userGUID string
	queryUser := fmt.Sprintf("SELECT login FROM %s WHERE login = $1", usersTable)
	rowUser := r.db.QueryRowContext(ctx, queryUser, user.Login)

	if err := rowUser.Scan(&existUser); err == nil {
		// newErr := NewUserError(err)
		return "", nil
	}
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	query := fmt.Sprintf("INSERT INTO %s (login, user_guid, password_hash) values ($1, $2, $3) RETURNING user_guid", usersTable)
	row := tx.QueryRowContext(ctx, query, user.Login, user.Login, user.Password)
	if err := row.Scan(&userGUID); err != nil {
		tx.Rollback()
		return "", err
	}
	query = fmt.Sprintf("INSERT INTO %s (user_guid, current, withdrawn) values ($1, $2, $3)", balanceTable)
	_, err = tx.ExecContext(ctx, query, user.Login, 0, 0)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}
	return userGUID, nil
}

func (r *AuthPostgres) GetUser(ctx context.Context, user app.User) (string, error) {
	var userGUID string
	query := fmt.Sprintf("SELECT user_guid  FROM %s WHERE login = $1 AND password_hash = $2", usersTable)
	err := r.db.GetContext(ctx, &userGUID, query, user.Login, user.Password)
	if err != nil {
		return "", err
	}
	return userGUID, nil
}
