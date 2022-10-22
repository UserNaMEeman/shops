package config

import (
	"flag"
	"os"
)

func flagParse() (*string, *string, *string) {
	addr := flag.String("a", "8080", `address to run HTTP server (default ":8080")`)
	dbAddr := flag.String("d", "", "URI to database")
	asAddr := flag.String("r", "", "accural system address")
	flag.Parsed()
	return addr, dbAddr, asAddr
}

func GetConfig() (string, string, string) {
	flagAddr, flagDbAddr, flagAsAddr := flagParse()
	addr := os.Getenv("RUN_ADDRESS")
	dbAddr := os.Getenv("DATABASE_URI")
	asAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if addr == "" {
		addr = *flagAddr
	}
	if dbAddr == "" {
		dbAddr = *flagDbAddr
	}
	if asAddr == "" {
		asAddr = *flagAsAddr
	}
	return addr, dbAddr, asAddr
}
