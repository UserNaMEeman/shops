package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	ordersTable  = "orders"
	usersTable   = "users"
	balanceTable = "balance"
	buysTable    = "buys"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// "postgres://postgres:password@localhost:5432/gophermarket?sslmode=disable"
// func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
func NewPostgresDB(URI string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", URI)
	// db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sqlx.DB) []error {
	var errors []error
	query := fmt.Sprintf(`CREATE TABLE %s
		(
			id serial not null unique,
			login varchar(255) not null unique,
			user_guid varchar(255) not null unique,
			password_hash varchar(255) not null
		)`, usersTable)
	if _, err := db.Exec(query); err != nil {
		errors = append(errors, err)
		// return err
	}
	query = fmt.Sprintf(`CREATE TABLE %s
	(
		id serial not null unique,
		user_guid varchar(255) not null,
		value varchar(255) not null,
		date timestamp not null
	)`, ordersTable)
	if _, err := db.Exec(query); err != nil {
		errors = append(errors, err)
		// fmt.Println(err)
	}
	query = fmt.Sprintf(`CREATE TABLE %s
	(
		id serial not null unique,
		user_guid varchar(255) not null,
		current numeric,
		withdrawn numeric
	)`, balanceTable)
	if _, err := db.Exec(query); err != nil {
		errors = append(errors, err)
	}
	query = fmt.Sprintf(`CREATE TABLE %s
	(
		id serial not null unique,
		order_num varchar(255) not null,
		accrual numeric default 0
	)`, "accrual")
	if _, err := db.Exec(query); err != nil {
		fmt.Println(err)
		errors = append(errors, err)
	}
	query = fmt.Sprintf(`CREATE TABLE %s
	(
		id serial not null unique,
		user_guid varchar(255) not null,
		order_buy varchar(255) not null,
		sum numeric default 0,
		date_buy timestamp not null
	)`, buysTable)
	if _, err := db.Exec(query); err != nil {
		fmt.Println(err)
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}
