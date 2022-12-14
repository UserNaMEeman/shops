package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

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

func (order *Order) UploadOrderNumber(ctx context.Context, guid, orderNumber string) error {
	num, err := strconv.Atoi(orderNumber)
	if err != nil {
		return err
	}
	if !valid(num) {
		return errors.New("number is not valid")
	}
	return order.repo.UploadOrderNumber(ctx, guid, orderNumber)
}

func (order *Order) CheckOrder(ctx context.Context, guid, orderNumber string) (string, bool) {
	return order.repo.CheckOrder(ctx, guid, orderNumber)
}

func (order *Order) CheckValidOrder(orderNumber string) (bool, error) {
	num, err := strconv.Atoi(orderNumber)
	if err != nil {
		return false, err
	}
	if valid(num) {
		return true, nil
	} else {
		return false, nil
	}
}

func (order *Order) GetOrders(ctx context.Context, guid string) ([]app.UserOrders, error) {
	orders, err := order.repo.GetOrders(ctx, guid)
	fmt.Println("orders for ", guid, ": ", orders)
	if err != nil {
		fmt.Println("get orders err for ", guid)
		return []app.UserOrders{}, err
	}
	// return orders, nil
	var respOrders []app.UserOrders
	for _, i := range orders {
		res, err := order.accrualOrder(i.Order)
		if err != nil {
			// fmt.Println(err)
			return []app.UserOrders{}, err
		}
		if res.Order == "" || res.Status == "" {
			// fmt.Println("res is empty")
			return []app.UserOrders{}, nil
			// continue
		}
		userOrder := app.UserOrders{
			Order:   res.Order,
			Status:  res.Status,
			Accrual: res.Accrual,
			Date:    i.Date,
		}
		respOrders = append(respOrders, userOrder)
	}
	return respOrders, nil
}

func (order *Order) accrualOrder(number string) (app.Accruals, error) {
	numberURL := order.asURL + "/api/orders/" + number
	res, err := getOrder(numberURL)
	if err != nil {
		// fmt.Println(err)
		return app.Accruals{}, err
	}
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
		fmt.Println("status code for ", url, ": ", resp.StatusCode)
		numb := strings.Split(url, "/api/orders/")[1]
		return app.Accruals{Order: numb, Status: "NEW"}, nil
	}
	if err != nil {
		return app.Accruals{}, err
	}
	data, err := io.ReadAll(body)
	if err != nil {
		return app.Accruals{}, err
	}
	fmt.Println("raw response for ", url, ": ", string(data))
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
