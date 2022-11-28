package repository

import (
	"context"
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

func (r *OrderPostgres) UploadOrderNumber(ctx context.Context, userGUID, orderNumber string) error {
	timeNow := time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("INSERT INTO %s (user_guid, value, date) values ($1, $2, $3)", ordersTable)
	_, err := r.db.ExecContext(ctx, query, userGUID, orderNumber, timeNow) //.QueryRow(query, userGUID, orderNumber)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderPostgres) CheckOrder(ctx context.Context, guid, orderNumber string) (string, bool) {
	var userGUID string
	queryOrder := fmt.Sprintf("SELECT user_guid FROM %s WHERE value = $1", ordersTable)
	row := r.db.QueryRowContext(ctx, queryOrder, orderNumber)
	if err := row.Scan(&userGUID); err != nil {
		// fmt.Println(err)
		return "", true
	}
	return userGUID, false
}

func (r *OrderPostgres) GetOrders(ctx context.Context, guid string) ([]app.UserOrders, error) {
	var order app.UserOrders
	var orders []app.UserOrders
	queryOrder := fmt.Sprintf("SELECT value, date FROM %s WHERE user_guid = $1 ORDER BY date", ordersTable)
	rows, err := r.db.QueryContext(ctx, queryOrder, guid) //(queryOrder, guid)
	if err != nil {
		return orders, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&order.Order, &order.Date); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		return orders, err
	}
	return orders, nil
}
