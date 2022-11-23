package service

import (
	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
)

type Balance struct {
	repo repository.AccountingUser
}

func NewBalanceService(repo repository.AccountingUser) *Balance {
	return &Balance{repo: repo}
}

func (b *Balance) GetBalance(guid string, totalAccrual float64) (app.Balance, error) {
	return b.repo.GetBalance(guid, totalAccrual)
}
