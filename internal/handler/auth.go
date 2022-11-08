package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/UserNaMEeman/shops/app"
)

// func signUpDepricated(w http.ResponseWriter, r *http.Request) {
// 	user := info.NewUser()
// 	data, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 	}
// 	defer r.Body.Close()
// 	json.Unmarshal(data, &user)
// 	if user.Login == "" || user.Password == "" {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	conn := storage.Connect()
// 	userNameAvail, err := storage.AddUser(conn, user)
// 	if !userNameAvail {
// 		w.WriteHeader(409)
// 		w.Write([]byte("Username already exist\n"))
// 		return
// 	}
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	fmt.Println("user login: ", user.Login, "user password: ", user.Password)
// 	cook, err := storage.GenCook(user.Login)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	cookie := &http.Cookie{
// 		Name:   user.Login,
// 		Value:  cook,
// 		MaxAge: 300,
// 	}
// 	http.SetCookie(w, cookie)
// 	w.WriteHeader(http.StatusOK)
// }
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
	id, err := h.services.Authorization.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("login already exist"))
		return
	}
	token, err := h.services.Authorization.GenerateToken(newUser)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}
	w.Header().Add("Authorization", token) //"Bearer "+token
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("do test")
	w.Write([]byte("OKs"))
	// w.WriteHeader(http.StatusOK)
}

// func (h *Handler) signInDep(w http.ResponseWriter, r *http.Request) {
// 	var user app.User
// 	data, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	defer r.Body.Close()
// 	json.Unmarshal(data, &user)
// 	if user.Login == "" || user.Password == "" {
// 		w.WriteHeader(400)
// 		w.Write([]byte("login and password must be not empty"))
// 		return
// 	}
// 	cookie := h.services.Authorization.GenerateCookie(user)
// 	http.SetCookie(w, cookie)
// 	w.WriteHeader(200)
// }

// func Test(w http.ResponseWriter, r *http.Request) {
// 	user := info.NewUser()
// 	username := strings.Split(r.URL.String(), "/")[2]
// 	user.Login = username
// 	conn := storage.Connect()
// 	userNameAvail, err := storage.AddUser(conn, user)
// 	if !userNameAvail {
// 		w.WriteHeader(409)
// 		w.Write([]byte("Username already exist\n"))
// 		return
// 	}
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }
