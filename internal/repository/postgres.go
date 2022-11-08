package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	ordersTable = "orders"
	usersTable  = "users"
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
