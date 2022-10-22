package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/UserNaMEeman/shops/internal/config"
	"github.com/UserNaMEeman/shops/internal/handler"
	"github.com/go-chi/chi"
)

var (
	addr   *string
	dbAddr *string
	asAddr *string
)

func init() {
	addr = flag.String("a", ":8080", `address to run HTTP server (default ":8080")`)
	dbAddr = flag.String("d", "", "URI to database")
	asAddr = flag.String("r", "", "accural system address")
}

func main() {
	flag.Parse()
	r := chi.NewRouter()
	addr, db, as := config.GetConfig(addr, dbAddr, asAddr)
	fmt.Println("programm parametrs: ", addr, db, as)
	r.Post("/api/user/register", handler.Register)

	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Println(err)
	}
}
