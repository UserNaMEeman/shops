package service

import (
	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
)

type Authorization interface {
	CreateUser(user app.User) (string, error)
	GenerateToken(user app.User) (string, error)
	ParseToken(token string) (string, error)
	// MatchGUIDTOKEN(guid, token string) error
}

type Orders interface {
	UploadOrderNumber(guid, number string) error
	CheckOrder(guid, orderNumber string) (string, bool)
	GetOrders(guid string) ([]app.UserOrders, error)
	// accrualOrder() app.Accruals
} //приём номеров заказов от зарегистрированных пользователей;

// type AccrualOrder interface {
// 	GetAccrualInformation(urlAccrualSystem string) (app.Accruals, error)
// }

type AccountingOrders interface {
} //учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;

type AccountingUser interface{} //учёт и ведение накопительного счёта зарегистрированного пользователя;

type LoyaltyPoints interface{} //проверка принятых номеров заказов через систему расчёта баллов лояльности;

type Rewards interface{} //начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

type Service struct {
	Authorization
	Orders
	// AccrualOrder
	AccountingOrders
	AccountingUser
	LoyaltyPoints
	Rewards
}

func NewServices(repos *repository.Repository, asURL string) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Orders:        NewOrdersService(repos.Orders, asURL),
		// AccrualOrder:  NewAccrualService(accrual.AccrualOrder),
	}
}
