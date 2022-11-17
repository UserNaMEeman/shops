package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
)

type Order struct {
	repo  repository.Orders
	asURL string
}

func NewOrdersService(repo repository.Orders, asURL string) *Order {
	return &Order{repo: repo, asURL: asURL}
}

// type GetAccrual struct {
// 	accrual accrual.AccrualOrder
// }

// func NewAccrualService(accrual accrual.AccrualOrder) *GetAccrual {
// 	return &GetAccrual{accrual: accrual}
// }

func (order *Order) accrualOrder(number string) app.Accruals {
	numberURL := order.asURL + "//api/orders/" + number
	res, err := occrualOrder(numberURL)
	if err != nil {
		fmt.Println(err)
	}
	// res := app.Accruals{}
	return res
}

func (order *Order) UploadOrderNumber(guid, orderNumber string) error {
	num, err := strconv.Atoi(orderNumber)
	if err != nil {
		return err
	}
	if !valid(num) {
		return errors.New("number is not valid")
	}
	return order.repo.UploadOrderNumber(guid, orderNumber)
}

func (order *Order) CheckOrder(guid, orderNumber string) (string, bool) {
	return order.repo.CheckOrder(guid, orderNumber)
}

func (order *Order) GetOrders(guid string) ([]app.UserOrders, error) {
	// return order.repo.GetOrders(guid)
	orders, err := order.repo.GetOrders(guid)
	if err != nil {
		return orders, err
	}
	for i := 0; i < len(orders); i++ {
		res := order.accrualOrder(orders[i].Order)
		orders[i].Accrual = res.Accrual
		orders[i].Status = res.Status
		// fmt.Println(&a[i])
	}
	return orders, nil
}

func valid(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

func occrualOrder(url string) (app.Accruals, error) {
	accrual := app.Accruals{}
	clinet := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
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
