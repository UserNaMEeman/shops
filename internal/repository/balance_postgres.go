package repository

import (
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
	// var withdrawn float64
	// balnce := app.Balance{}
	// queryOrder := fmt.Sprintf("SELECT withdrawn FROM %s WHERE user_guid = $1", balanceTable)
	// row := r.db.QueryRow(queryOrder, guid) //(queryOrder, guid)
	// if err := row.Scan(&withdrawn); err != nil {
	// 	return app.Balance{}, err
	// }
	// balnce.Current = totalAccrual - withdrawn
	// balnce.Withdrawn = withdrawn
	// fmt.Println("postgre balance: ", balnce)
	// return balnce, nil
	a := app.Balance{totalAccrual, 0}
	return a, nil
}
