package main

import (
	"flag"
	"net/http"

	"github.com/UserNaMEeman/shops/internal/config"
	"github.com/UserNaMEeman/shops/internal/handler"
	"github.com/UserNaMEeman/shops/internal/repository"
	"github.com/UserNaMEeman/shops/internal/service"
	"github.com/sirupsen/logrus"
	// "github.com/UserNaMEeman/shops/internal/storage"
)

var (
	addr   *string
	dbAddr *string
	asAddr *string
)

func init() {
	addr = flag.String("a", "localhost:8080", `address to run HTTP server (default ":8080")`)
	dbAddr = flag.String("d", "", "URI to database")
	asAddr = flag.String("r", "", "accural system address")
}

//  ***postgres/praktikum?sslmode=disable
func main() {
	// logrus.SetFormatter(new(logrus.JSONFormatter))
	flag.Parse()
	addr, dbURI, as := config.GetConfig(addr, dbAddr, asAddr)
	db, err := repository.NewPostgresDB(dbURI)
	if err != nil {
		// fmt.Println(err)
		logrus.Fatalf("failed connect to DB: %s", err.Error())
	}
	if errs := repository.CreateTables(db); err != nil {
		logrus.Println(errs)
	}
	repos := repository.NewRepository(db)
	services := service.NewServices(repos, as)
	handlers := handler.NewHandler(services)
	// srv := new(app.Server)
	if err := http.ListenAndServe(addr, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
