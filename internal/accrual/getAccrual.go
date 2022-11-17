package accrual

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/UserNaMEeman/shops/app"
)

type GetAccrual struct {
	url string
}

func NewAccrualCreate(url string) *GetAccrual {
	return &GetAccrual{url: url}
}

func (a *GetAccrual) GetAccrualInformation(urlAccrualSystem string) (app.Accruals, error) {
	accrual := app.Accruals{}
	clinet := http.Client{}
	request, err := http.NewRequest("GET", urlAccrualSystem, nil)
	if err != nil {
		return accrual, err
	}
	resp, err := clinet.Do(request)
	if err != nil {
		return accrual, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return accrual, err
	}
	json.Unmarshal(data, &accrual)
	return accrual, nil
}
