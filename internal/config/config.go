package config

import (
	"os"
)

func GetConfig(flagAddr, flagDBAddr, flagAsAddr *string) (string, string, string) {
	addr := os.Getenv("RUN_ADDRESS")
	dbAddr := os.Getenv("DATABASE_URI")
	asAddr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if addr == "" {
		addr = *flagAddr
	}
	if dbAddr == "" {
		dbAddr = *flagDBAddr
	}
	if asAddr == "" {
		asAddr = *flagAsAddr
	}
	return addr, dbAddr, asAddr
}
