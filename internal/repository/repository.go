package repository

import (
	"github.com/UserNaMEeman/shops/app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface { //регистрация, аутентификация и авторизация пользователей;
	CreateUser(app.User) (int, error)
	GetUser(app.User) (string, error)
}

type Orders interface {
	GetNumber(app.User) (string, error)
} //приём номеров заказов от зарегистрированных пользователей;

type AccountingOrders interface{} //учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;

type AccountingUser interface{} //учёт и ведение накопительного счёта зарегистрированного пользователя;

type LoyaltyPoints interface{} //проверка принятых номеров заказов через систему расчёта баллов лояльности;

type Rewards interface{} //начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

type Repository struct {
	Authorization
	Orders
	AccountingOrders
	AccountingUser
	LoyaltyPoints
	Rewards
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
