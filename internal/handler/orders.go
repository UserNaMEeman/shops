package handler

import (
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) uploadOrder(w http.ResponseWriter, r *http.Request) {
	// var order
	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	guid := fmt.Sprintf("%s", ctx.Value("guid"))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order := h.services.Orders
	orderGUID, free := order.CheckOrder(guid, string(body))
	if free {
		if err = order.UploadOrderNumber(guid, string(body)); err != nil {
			// fmt.Println(err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}
	if orderGUID == guid {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// h.services.UploadOrderNumber(string(order))
}
