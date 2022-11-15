package service

import (
	"crypto/md5"
	"encoding/hex"

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
} //приём номеров заказов от зарегистрированных пользователей;

type AccountingOrders interface{} //учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;

type AccountingUser interface{} //учёт и ведение накопительного счёта зарегистрированного пользователя;

type LoyaltyPoints interface{} //проверка принятых номеров заказов через систему расчёта баллов лояльности;

type Rewards interface{} //начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.

type Service struct {
	Authorization
	Orders
	AccountingOrders
	AccountingUser
	LoyaltyPoints
	Rewards
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Orders:        NewOrdersService(repos.Orders),
	}
}

////////////////////////////
///////////////////////////
//////////////////////////
/////////////////////////
////////////////////////
///////////////////////
//////////////////////
/////////////////////
////////////////////
///////////////////
//////////////////
/////////////////
////////////////
///////////////
//////////////
/////////////
////////////
///////////
//////////
/////////
////////
///////
//////
/////
////
///
//
type Hash interface {
	GenerateHash(string) string
	// CheckHash(string) bool
}

type MyHash struct{}

func (*MyHash) GenerateHash(password string) string {
	hash := md5.Sum([]byte(password))
	// return string(hash[:])
	return hex.EncodeToString(hash[:])
}

// func (*md5Hash) CheckHash(password string) string {
// 	hash := md5.Sum([]byte(password))
// 	return string(hash[:])
// }
