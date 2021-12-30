package main

import (
	"fmt"
	"os"

	"github.com/SohailHaqyar/wait-so/database"
	"github.com/SohailHaqyar/wait-so/notes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	URI := os.Getenv("DATABASE_URL")

	database.DatabaseConfig, err = gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Error while connecting to the database")
	}
	fmt.Println("Connected to the database successfully")
	database.DatabaseConfig.AutoMigrate(notes.Note{})
}

func setupRoutes(app *fiber.App) {
	notes.SetupRoutes(app)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	port := os.Getenv("PORT")

	fmt.Println("Starting the server on port 8000")

	initDatabase()
	setupRoutes(app)

	app.Listen(":" + port)
	sqldb, _ := database.DatabaseConfig.DB()
	defer sqldb.Close()
}
