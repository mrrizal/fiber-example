package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mrrizal/fiber-example/book"
	"github.com/mrrizal/fiber-example/routes"
	log "github.com/sirupsen/logrus"

	"github.com/mrrizal/fiber-example/configs"
	"github.com/mrrizal/fiber-example/database"
)

func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Info("Failed load .env file")
	}
	log.Info("Successful load .env file")
}

func main() {
	loadEnvFile()

	configs.GetSettings()
	database.InitDatabase()

	migrate := flag.Bool("migrate", false, "a bool")
	flag.Parse()

	if *migrate {
		database.DBConn.AutoMigrate(&book.Book{})
		log.Info("Database migrated.")
	} else {
		app := fiber.New()
		routes.SetupRoutes(app)
		app.Listen("0.0.0.0:3000")
	}
}
