package main

import (
	"flag"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mrrizal/fiber-example/book"
	"github.com/mrrizal/fiber-example/routes"
	"github.com/mrrizal/fiber-example/user"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmfiber"

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

func setup(bookService book.Service, userService user.Service) *fiber.App {
	app := fiber.New()
	app.Use(apmfiber.Middleware())
	routes.SetupRoutes(app, bookService, userService)
	return app
}

func main() {
	loadEnvFile()

	configs.GetSettings()
	time.Sleep(3 * time.Second)
	database.InitDatabase()

	migrate := flag.Bool("migrate", false, "a bool")
	flag.Parse()

	if *migrate {
		database.DBConn.AutoMigrate(&user.User{}, &book.Book{})
		log.Info("Database migrated.")
	} else {
		app := setup(&book.ServiceStruct{}, &user.ServiceStruct{})
		app.Listen("0.0.0.0:3000")
	}
}
