package handler

import (
	"fmt"
	"net/http"
)

// func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
// 	authHeader := r.Header.Get("")
// }

// func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
// 	tokenCookie, err := r.Cookie("Customer3")
// 	if err != nil {
// 		log.Fatalf("Error occured while reading cookie")
// 	}
// 	fmt.Println(tokenCookie)
// }

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, err := h.services.Authorization.ParseToken(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
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
