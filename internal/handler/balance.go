package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UserNaMEeman/shops/app"
)

func (h *Handler) getTotalBalance(guid string) (float64, error) {
	var totalBalace float64
	newOrder := h.services.Orders
	orders, err := newOrder.GetOrders(guid)
	if err != nil {
		// fmt.Println("GetOrders err: ", err)
		return 0, err
	}
	for _, i := range orders {
		totalBalace = totalBalace + i.Accrual
		// fmt.Println("order current: ", i.Accrual)
	}
	return totalBalace, nil
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	guid := fmt.Sprintf("%s", ctx.Value("guid"))
	totalBalace, err := h.getTotalBalance(guid)
	if err != nil {
		// fmt.Println("GetOrders err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("totalBalance: ", totalBalace)
	newBalance := h.services.AccountingUser
	// fmt.Println("newBalance: ", newBalance)
	balance, err := newBalance.GetBalance(guid, totalBalace)
	if err != nil {
		// fmt.Println("balance err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(balance)
	if err != nil {
		// fmt.Println("Marshal err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusInternalServerError)
	}
	ctx := r.Context()
	guid := fmt.Sprintf("%s", ctx.Value("guid"))
	totalBalace, err := h.getTotalBalance(guid)
	if err != nil {
		// fmt.Println("GetOrders err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newBalance := h.services.AccountingUser
	newOrder := h.services.Orders
	balance, err := newBalance.GetBalance(guid, totalBalace)
	if err != nil {
		// fmt.Println("balance err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buy := app.Buy{}
	body := r.Body
	defer r.Body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		// fmt.Println("body err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.Unmarshal(data, &buy)
	if balance.Current < buy.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	valid, err := newOrder.CheckValidOrder(buy.Order)
	if err != nil {
		// fmt.Println("valid err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !valid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if err = newBalance.UsePoints(guid, buy); err != nil {
		// fmt.Println("Use Points: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("guid: ", guid, "order: ", buy.Order)
	w.WriteHeader(http.StatusOK)
}
