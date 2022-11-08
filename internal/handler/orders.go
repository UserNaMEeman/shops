package handler

import (
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) uploadOrder(w http.ResponseWriter, r *http.Request) {
	order, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(order))
}
