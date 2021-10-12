package configs

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type configs struct {
	DatabaseURL string
	PageSize    int
	SecretKey   string
}

var Configs configs

func GetSettings() {
	Configs.DatabaseURL = os.Getenv("DATABASE_URL")
	pageSize, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))
	if err != nil {
		log.Panic(err)
	}
	Configs.PageSize = pageSize
	Configs.SecretKey = os.Getenv("SECRET_KEY")
}
