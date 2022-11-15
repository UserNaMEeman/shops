package handler

import (
	"context"
	"fmt"
	"net/http"
)

// type guid string

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	// fmt.Println("//////////////////////userIdentity////////////////////////////////")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// fmt.Println("authHeader: ", authHeader)
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userGUID, err := h.services.Authorization.ParseToken(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// userGUID := "test"
		// fmt.Println("//////////////////////Context////////////////////////////////")
		ctx := r.Context()
		// fmt.Println("//////////////////////ctx////////////////////////////////")
		ctx = context.WithValue(ctx, "guid", userGUID)
		// fmt.Println("//////////////////////guid////////////////////////////////")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) IsLoggedIn(w http.ResponseWriter, r *http.Request) {

	// Obtains cookie from users http.Request
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		// return false/
	}
	sessionID := cookie.Value
	fmt.Println(sessionID)
}
