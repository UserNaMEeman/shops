package main

import (
	"fmt"

	"github.com/UserNaMEeman/shops/internal/config"
)

func main() {
	addr, db, as := config.GetConfig()
	fmt.Println("addr: ", addr, " db: ", db, " as: ", as)
}
