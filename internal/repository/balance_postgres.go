package repository

import (
	"fmt"
	"time"

	"github.com/UserNaMEeman/shops/app"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetBalance(guid string, totalAccrual float64) (app.Balance, error) {
	var withdrawn float64
	balnce := app.Balance{}
	queryOrder := fmt.Sprintf("SELECT withdrawn FROM %s WHERE user_guid = $1", balanceTable)
	row := r.db.QueryRow(queryOrder, guid) //(queryOrder, guid)
	if err := row.Scan(&withdrawn); err != nil {
		return app.Balance{}, err
	}
	balnce.Current = totalAccrual - withdrawn
	balnce.Withdrawn = withdrawn
	// fmt.Println("postgre balance: ", balnce)
	return balnce, nil
}

func (r *BalancePostgres) UsePoints(guid string, buy app.Buy) error { //add update orders table
	timeNow := buy.Date.Format(time.RFC3339)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_guid, order_buy, sum, date_buy) values ($1, $2, $3, $4)", buysTable)
	_, err = tx.Exec(query, guid, buy.Order, buy.Sum, timeNow) //.QueryRow(query, userGUID, orderNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("UPDATE %s SET withdrawn = withdrawn + $1 WHERE user_guid = $2", balanceTable)
	_, err = tx.Exec(query, buy.Sum, guid)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	// }
	return nil
}

func (r *BalancePostgres) GetWithdrawals(guid string) ([]app.Buy, error) {
	var withdrawn app.Buy
	var withdrawals []app.Buy
	queryOrder := fmt.Sprintf("SELECT order_buy, sum, date_buy FROM %s WHERE user_guid = $1 ORDER BY date_buy", buysTable)
	rows, err := r.db.Query(queryOrder, guid) //(queryOrder, guid)
	if err != nil {
		return withdrawals, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&withdrawn.Order, &withdrawn.Sum, &withdrawn.Date); err != nil {
			return withdrawals, err
		}
		// fmt.Println("order: ", order)
		withdrawals = append(withdrawals, withdrawn)
	}
	err = rows.Err()
	if err != nil {
		return withdrawals, err
	}
	return withdrawals, nil
}
