package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/UserNaMEeman/shops/internal/info"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := info.NewUser()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	defer r.Body.Close()
	json.Unmarshal(data, &user)
	fmt.Println(user.Login, user.Password)
	w.Header().Add("Authorization", "ok")
	// w.WriteHeader("Authorization")
	w.WriteHeader(http.StatusOK)
}
