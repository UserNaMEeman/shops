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
		next.ServeHTTP(w, r)
	})
	// tokenCookie, err := r.Cookie("Customer3")
	// if err != nil {
	// 	log.Fatalf("Error occured while reading cookie")
	// }
	// fmt.Println(tokenCookie)
}

func (h *Handler) IsLoggedIn(w http.ResponseWriter, r *http.Request) {

	// Obtains cookie from users http.Request
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		// return false/
	}

	// Obtain sessionID from cookie we obtained earlier
	sessionID := cookie.Value
	fmt.Println(sessionID)
	// Split the sessionID to Username and ID (username+random)
	// z := strings.Split(sessionID, ":")
	// email := z[0]
	// sessionID = z[1]

	// // If SessionID matches the expected SessionID, it is Good
	// if sessionID == expectedSessionID {
	// 	// If you want to be really secure check IP
	// 	return true
	// }

	// return false
}
