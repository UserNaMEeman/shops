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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, i := range orders {
		totalBalace = totalBalace + i.Accrual
	}
	newBalance := h.services.AccountingUser
	balance, err := newBalance.GetBalance(guid, totalBalace)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
