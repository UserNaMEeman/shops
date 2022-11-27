package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/UserNaMEeman/shops/app"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var newUser app.User
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()
	json.Unmarshal(data, &newUser)
	if newUser.Login == "" || newUser.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte("login and password must be not empty"))
		return
	}
	GUID, err := h.services.Authorization.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if GUID == "" {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("login already exist"))
		return
	}
	token, err := h.services.Authorization.GenerateToken(newUser)
	if err != nil {
		fmt.Println("gen token err: ", err)
		w.WriteHeader(401)
		return
	}
	w.Header().Add("Authorization", token) //"Bearer "+token
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) { // must add 401 — неверная пара логин/пароль;
	var user app.User
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	json.Unmarshal(data, &user)
	if user.Login == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login and password must be not empty"))
		return
	}
	token, err := h.services.Authorization.GenerateToken(user)
	if err != nil {
		fmt.Println("gen token err: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Add("Authorization", token) //"Bearer "+token
	w.WriteHeader(http.StatusOK)
}
