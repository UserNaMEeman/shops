package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/UserNaMEeman/shops/app"
)

func (h *Handler) getTotalBalance(ctx context.Context, guid string) (float64, error) {
	var totalBalace float64
	newOrder := h.services.Orders
	orders, err := newOrder.GetOrders(ctx, guid)
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

func (h *Handler) getBalance(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	guid := fmt.Sprintf("%s", ctx.Value(app.TypeGUID))
	totalBalace, err := h.getTotalBalance(ctx, guid)
	if err != nil {
		// fmt.Println("GetOrders err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("totalBalance: ", totalBalace)
	newBalance := h.services.AccountingUser
	// fmt.Println("newBalance: ", newBalance)
	balance, err := newBalance.GetBalance(ctx, guid, totalBalace)
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

func (h *Handler) withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// ctx := r.Context()
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	guid := fmt.Sprintf("%s", ctx.Value(app.TypeGUID))
	totalBalace, err := h.getTotalBalance(ctx, guid)
	if err != nil {
		// fmt.Println("GetOrders err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newBalance := h.services.AccountingUser
	newOrder := h.services.Orders
	balance, err := newBalance.GetBalance(ctx, guid, totalBalace)
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
	if err = newBalance.UsePoints(ctx, guid, buy); err != nil {
		// fmt.Println("Use Points: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("guid: ", guid, "order: ", buy.Order)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) withdrawals(w http.ResponseWriter, r *http.Request) {
	newBalance := h.services.AccountingUser
	// ctx := r.Context()
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	guid := fmt.Sprintf("%s", ctx.Value(app.TypeGUID))

	buys, err := newBalance.GetWithdrawals(ctx, guid)
	if err != nil {
		if len(buys) == 0 {
			// fmt.Println("valid err: ", err)
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	data, err := json.Marshal(buys)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("send wths: ", string(data))
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
