package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/routes"
)

func main() {

	godotenv.Load(".env")

	database.Connect()

	port := os.Getenv("port")
	allowed_origins := os.Getenv("allowed_origins")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowed_origins,
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%s", port))
}
