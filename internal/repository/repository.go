package repository

import (
	"context"

	"github.com/UserNaMEeman/shops/app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface { //регистрация, аутентификация и авторизация пользователей;
	CreateUser(context.Context, app.User) (string, error)
	GetUser(context.Context, app.User) (string, error)
}

type Orders interface {
	UploadOrderNumber(ctx context.Context, guid, orderNumber string) error
	CheckOrder(ctx context.Context, guid, orderNumber string) (string, bool)
	GetOrders(ctx context.Context, guid string) ([]app.UserOrders, error)
} //приём номеров заказов от зарегистрированных пользователей;

type AccountingOrders interface{} //учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;

type AccountingUser interface {
	GetBalance(ctx context.Context, guid string, totalAccrual float64) (app.Balance, error)
	UsePoints(ctx context.Context, guid string, buy app.Buy) error
	GetWithdrawals(ctx context.Context, guid string) ([]app.Buy, error)
} //учёт и ведение накопительного счёта зарегистрированного пользователя;

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
		Authorization:  NewAuthPostgres(db),
		Orders:         NewOrderPostgres(db),
		AccountingUser: NewBalancePostgres(db),
	}
}
