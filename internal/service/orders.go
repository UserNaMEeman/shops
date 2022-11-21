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
	var respOrders []app.UserOrders
	for _, i := range orders {
		res, err := order.accrualOrder(i.Order)
		if err != nil {
			fmt.Println(err)
			return []app.UserOrders{}, err
		}
		if res.Order == "" || res.Status == "" {
			continue
		}
		// fmt.Println("accrual: ", res)
		respOrders = append(respOrders, i)
		fmt.Println("respOrders: ", respOrders)
	}
	return respOrders, nil
	// for i := 0; i < len(orders); i++ {
	// res, err := order.accrualOrder(orders[i].Order)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return []app.UserOrders{}, err
	// }
	// if res.Order == "" {

	// }
	// fmt.Println("accrual: ", res)
	// if res.Accrual != "" {
	// 	orders[i].Accrual = res.Accrual
	// }
	// if res.Status != "" {
	// 	orders[i].Status = res.Status
	// }
	// }
	// return respOrders, nil
}

func (order *Order) accrualOrder(number string) (app.Accruals, error) {
	numberURL := order.asURL + "/api/orders/" + number
	res, err := getOrder(numberURL)
	if err != nil {
		fmt.Println(err)
		return app.Accruals{}, err
	}
	// fmt.Println("accrual: ", res)
	return res, nil
}

func getOrder(url string) (app.Accruals, error) {
	accrual := app.Accruals{}
	clinet := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return app.Accruals{}, err
	}
	resp, err := clinet.Do(request)
	body := resp.Body
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("status code: ", resp.StatusCode)
		return app.Accruals{}, nil
	}
	if err != nil {
		return app.Accruals{}, err
	}
	data, err := io.ReadAll(body)
	if err != nil {
		return app.Accruals{}, err
	}
	json.Unmarshal(data, &accrual)
	return accrual, nil
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
