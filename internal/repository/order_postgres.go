package repository

import (
	"fmt"
	"time"

	"github.com/UserNaMEeman/shops/app"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) UploadOrderNumber(userGUID, orderNumber string) error {
	timeNow := time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO %s (user_guid, value, date) values ($1, $2, $3)", ordersTable)
	_, err := r.db.Exec(query, userGUID, orderNumber, timeNow) //.QueryRow(query, userGUID, orderNumber)

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

func (r *OrderPostgres) GetOrders(guid string) ([]app.UserOrders, error) {
	var order app.UserOrders
	var orders []app.UserOrders
	queryOrder := fmt.Sprintf("SELECT value, date FROM %s WHERE user_guid = $1 ORDER BY data", ordersTable)
	rows, err := r.db.Query(queryOrder, guid) //(queryOrder, guid)
	if err != nil {
		// fmt.Println(err)
		return orders, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&order.Order, &order.Date); err != nil {
			return orders, err
		}
		fmt.Println("order: ", order)
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		// fmt.Println(err)
		return orders, err
	}
	// fmt.Println("All orders: ", order)
	return orders, nil
}
