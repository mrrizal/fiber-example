package database

import (
	"errors"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

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

	errConnect := errors.New("cant connect")
	for i := 0; i < 5; i++ {
		DBConn, err = gorm.Open(postgres.Open(u))
		if err == nil {
			errConnect = nil
			break
		}
		errConnect = err
		time.Sleep(1 * time.Second)
	}

	if errConnect != nil {
		log.Panic(err.Error())
	}
	log.Info("Successful connected to database.")
}
