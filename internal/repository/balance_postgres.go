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
	fmt.Println("postgre balance: ", balnce)
	return balnce, nil
}

func (r *BalancePostgres) UsePoints(guid string, buy app.Buy) error { //add update orders table
	timeNow := buy.Date.Format(time.RFC3339)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (order_buy, sum, date_buy) values ($1, $2, $3)", buysTable)
	_, err = tx.Exec(query, buy.Order, buy.Sum, timeNow) //.QueryRow(query, userGUID, orderNumber)
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
	// query = fmt.Sprintf("INSERT INTO %s (user_guid, value, date) values ($1, $2, $3)", ordersTable)
	// _, err = r.db.Exec(query, guid, buy.Order, timeNow) //.QueryRow(query, userGUID, orderNumber)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	var g string
	queryOrder := fmt.Sprintf("SELECT user_guid FROM %s WHERE value = $1", ordersTable)
	row := r.db.QueryRow(queryOrder, buy.Order) //(queryOrder, guid)
	if err := row.Scan(&g); err != nil {
		fmt.Println("error get value guid: ", err)
	} else {
		fmt.Println("value guid: ", g)
	}
	return nil
}

func (r *BalancePostgres) GetWithdrawals(guid string) (app.Buy, error) {
	return app.Buy{}, nil
}
