package main

import (
	"flag"
	"fmt"

	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/config"
	"github.com/UserNaMEeman/shops/internal/handler"
	"github.com/UserNaMEeman/shops/internal/repository"
	"github.com/UserNaMEeman/shops/internal/service"

	// "github.com/UserNaMEeman/shops/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	addr   *string
	dbAddr *string
	asAddr *string
)

func init() {
	addr = flag.String("a", "8080", `address to run HTTP server (default ":8080")`)
	dbAddr = flag.String("d", "", "URI to database")
	asAddr = flag.String("r", "", "accural system address")
}

//  ***postgres/praktikum?sslmode=disable
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	flag.Parse()
	addr, dbURI, as := config.GetConfig(addr, dbAddr, asAddr)
	fmt.Println("programm parametrs: ", addr, dbURI, as)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "password",
		DBName:   "gophermarket",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed connect to DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)
	srv := new(app.Server)
	if err := srv.Run(addr, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
	// storage.Connect()
}
