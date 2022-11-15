package service

import (
	"errors"
	"strconv"

	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
)

type Order struct {
	repo repository.Orders
}

func NewOrdersService(repo repository.Orders) *Order {
	return &Order{repo: repo}
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
	return order.repo.GetOrders(guid)
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
