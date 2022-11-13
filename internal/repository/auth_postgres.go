package repository

import (
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

func (r *AuthPostgres) CreateUser(user app.User) (string, error) {
	var existUser string
	var userGUID string
	queryUser := fmt.Sprintf("SELECT login FROM %s WHERE login = $1", usersTable)
	rowUser := r.db.QueryRow(queryUser, user.Login)
	if err := rowUser.Scan(&existUser); err == nil {
		// newErr := NewUserError(err)
		return "", nil
	}
	query := fmt.Sprintf("INSERT INTO %s (login, user_guid, password_hash) values ($1, $2, $3) RETURNING user_guid", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Login, user.Password)
	if err := row.Scan(&userGUID); err != nil {
		return "", err
	}
	return userGUID, nil
}

func (r *AuthPostgres) GetUser(user app.User) (string, error) {
	var userGUID string
	query := fmt.Sprintf("SELECT user_guid  FROM %s WHERE login = $1 AND password_hash = $2", usersTable)
	err := r.db.Get(&userGUID, query, user.Login, user.Password)
	if err != nil {
		return "", err
	}
	return userGUID, nil
}

// type usererr struct {
// 	message string
// 	err     error
// }

// func (mes *usererr) Error() string {
// 	return fmt.Sprintf("%s %v", mes.message, mes.err)
// }

// func NewUserError(err error) error {
// 	return &usererr{
// 		message: "user:",
// 		err:     errors.New("already exist"),
// 	}
// }
