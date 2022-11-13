package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) UploadOrderNumber(userGUID, orderNumber string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_guid, value) values ($1, $2)", ordersTable)
	_, err := r.db.Exec(query, userGUID, orderNumber) //.QueryRow(query, userGUID, orderNumber)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderPostgres) CheckOrder(guid, orderNumber string) (string, bool) {
	var userGUID string
	queryOrder := fmt.Sprintf("SELECT user_guid FROM %s WHERE value = $1", ordersTable)
	row := r.db.QueryRow(queryOrder, orderNumber)
	if err := row.Scan(&userGUID); err != nil {
		// fmt.Println(err)
		return "", true
	}
	return userGUID, false
}
