package service

import (
	"context"

	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user app.User) (string, error)
	GenerateToken(ctx context.Context, user app.User) (string, error)
	ParseToken(token string) (string, error)
}

type Orders interface {
	UploadOrderNumber(ctx context.Context, guid, number string) error
	CheckOrder(ctx context.Context, guid, orderNumber string) (string, bool)
	GetOrders(ctx context.Context, guid string) ([]app.UserOrders, error)
	CheckValidOrder(orderNumber string) (bool, error)
} //приём номеров заказов от зарегистрированных пользователей;

// type AccountingOrders interface {
// } //учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;

type AccountingUser interface {
	GetBalance(ctx context.Context, guid string, totalAccrual float64) (app.Balance, error)
	UsePoints(ctx context.Context, guid string, buy app.Buy) error
	GetWithdrawals(ctx context.Context, guid string) ([]app.Buy, error)
} //учёт и ведение накопительного счёта зарегистрированного пользователя;

type WithdrawPoints interface {
} //проверка принятых номеров заказов через систему расчёта баллов лояльности;

type Rewards interface{} //начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

type Service struct {
	Authorization
	Orders
	AccountingUser
	WithdrawPoints
	Rewards
	// AccountingOrders
}

func NewServices(repos *repository.Repository, asURL string) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		Orders:         NewOrdersService(repos.Orders, asURL),
		AccountingUser: NewBalanceService(repos.AccountingUser),
	}
}
