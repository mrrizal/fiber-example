package database

import (
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mrrizal/fiber-example/configs"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	u, err := pq.ParseURL(configs.Configs.DatabaseURL)

	if err != nil {
		log.Panic(err.Error())
	}

	DBConn, err = gorm.Open(postgres.Open(u))
	if err != nil {
		log.Panic(err.Error())
	}
	log.Info("Successful connected to database.")
}
