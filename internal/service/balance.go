package service

import (
	"context"
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

func (b *Balance) GetBalance(ctx context.Context, guid string, totalAccrual float64) (app.Balance, error) {
	return b.repo.GetBalance(ctx, guid, totalAccrual)
}

func (b *Balance) UsePoints(ctx context.Context, guid string, buy app.Buy) error {
	buy.Date = time.Now()
	return b.repo.UsePoints(ctx, guid, buy)
}

func (b *Balance) GetWithdrawals(ctx context.Context, guid string) ([]app.Buy, error) {

	return b.repo.GetWithdrawals(ctx, guid)
}
