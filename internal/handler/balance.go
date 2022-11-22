package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var totalBalace float64
	ctx := r.Context()
	guid := fmt.Sprintf("%s", ctx.Value("guid"))
	newOrder := h.services.Orders
	orders, err := newOrder.GetOrders(guid)
	if err != nil {
		fmt.Println("GetOrders err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, i := range orders {
		totalBalace = totalBalace + i.Accrual
		fmt.Println("order current: ", i.Accrual)
	}
	fmt.Println("totalBalance: ", totalBalace)
	newBalance := h.services.AccountingUser
	fmt.Println("newBalance: ", newBalance)
	balance, err := newBalance.GetBalance(guid, totalBalace)
	if err != nil {
		fmt.Println("balance err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(balance)
	if err != nil {
		fmt.Println("Marshal err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
