package service

import (
	"time"

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

func (b *Balance) UsePoints(guid string, buy app.Buy) error {
	buy.Date = time.Now()
	return b.repo.UsePoints(guid, buy)
}

func (b *Balance) GetWithdrawals(guid string) (app.Buy, error) {

	return app.Buy{}, nil
}
