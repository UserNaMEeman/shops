package accrual

import (
	"github.com/UserNaMEeman/shops/app"
)

type AccrualOrder interface {
	GetAccrualInformation(urlAccrualSystem string) (app.Accruals, error)
}

type Accrual struct {
	AccrualOrder
}

func NewAccrual(url string) *Accrual {
	return &Accrual{
		AccrualOrder: NewAccrualCreate(url),
	}
}
