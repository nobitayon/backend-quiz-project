package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	controller "github.com/nobitayon/QuizProject/controllers"
	middleware_jwt "github.com/nobitayon/QuizProject/middleware/middleware_jwt"
	middleware_timeRequest "github.com/nobitayon/QuizProject/middleware/middleware_timeRequest"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/register", controller.Register)
	app.Post("/api/v1/register_admin", middleware_jwt.New(middleware_jwt.Config{Secret: "secretKey"}), controller.RegisterAdmin)
	app.Post("/api/v1/login", controller.Login)

	quizGroup := app.Group("/api/v1/quiz")
	quizGroup.Use(middleware_jwt.New(middleware_jwt.Config{Secret: "secretKey"}))
	quizGroup.Get("/", controller.GetAllQuiz)
	quizGroup.Get("/:idQuiz", controller.GetQuiz)
	quizGroup.Post("/", controller.AddQuiz)
	quizGroup.Patch("/:idQuiz", controller.EditQuiz)
	quizGroup.Delete("/:idQuiz", controller.DeleteQuiz)

	// app.Get("/api/quiz", controller.GetAllQuiz)
	// app.Get("/api/quiz/:idQuiz", controller.GetQuiz)
	// app.Post("/api/quiz", controller.AddQuiz)
	// app.Patch("/api/quiz/:idQuiz", controller.EditQuiz)
	// app.Delete("/api/quiz/:idQuiz", controller.DeleteQuiz)

	// pertanyaanGroup := app.Group("/api/v1/quiz")
	quizGroup.Post("/:idQuiz/pertanyaan", controller.CreatePertanyaanOnQuiz)
	quizGroup.Patch("/:idQuiz/pertanyaan/:idPertanyaan", controller.EditPertanyaanOnQuiz)
	quizGroup.Delete("/:idQuiz/pertanyaan/:idPertanyaan", controller.DeletePertanyaanOnQuiz)
	quizGroup.Get("/:idQuiz/pertanyaan/:idPertanyaan", controller.GetPertanyaanOnQuiz)
	quizGroup.Get("/:idQuiz/pertanyaan", controller.GetAllPertanyaanOnQuiz)

	// app.Post("/api/quiz/:idQuiz/pertanyaan", controller.CreatePertanyaanOnQuiz)
	// app.Patch("/api/quiz/:idQuiz/pertanyaan/:idPertanyaan", controller.EditPertanyaanOnQuiz)
	// app.Delete("/api/quiz/:idQuiz/pertanyaan/:idPertanyaan", controller.DeletePertanyaanOnQuiz)
	// app.Get("/api/quiz/:idQuiz/pertanyaan/:idPertanyaan", controller.GetPertanyaanOnQuiz)
	// app.Get("/api/quiz/:idQuiz/pertanyaan", controller.GetAllPertanyaanOnQuiz)

	app.Get("/api/v1/jawaban", controller.GetJawaban)
	app.Get("/api/v1/jawaban/:id", controller.GetJawabanById)
	app.Post("/api/v1/jawaban", controller.AddJawaban)
	app.Patch("/api/v1/jawaban/:id", controller.EditJawaban)
	app.Delete("/api/v1/jawaban/:id", controller.DeleteJawabanById)

	// cek restricted
	app.Get("/restricted", middleware_jwt.New(middleware_jwt.Config{Secret: "secretKey"}), func(c *fiber.Ctx) error {
		claims := c.Locals("jwtClaims").(jwt.MapClaims)

		return c.JSON(fiber.Map{"role": claims["role"]})

	})

	// cek time record
	app.Get("/recordTime", middleware_timeRequest.New(middleware_timeRequest.Config{}), func(c *fiber.Ctx) error {
		timeRecord := c.Locals("timeRecorded")
		return c.JSON(fiber.Map{"time": timeRecord})
	})

	// cek periksa
	app.Post("/periksa", middleware_jwt.New(middleware_jwt.Config{Secret: "secretKey"}), controller.PeriksaQuiz)
}
