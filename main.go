package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/routes"
)

func main() {

	godotenv.Load(".env")

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
