package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/models"
)

type inputPeriksaQuiz struct {
	IdQuiz int `json:"idQuiz"`
	IdUser int `json:"idUser"`
}

func PeriksaQuiz(c *fiber.Ctx) error {
	var data inputPeriksaQuiz

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	var quiz models.Quiz

	database.GormDB.Where("id = ?", data.IdQuiz).First(&quiz)

	if quiz.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "quiz yang ingin diperiksa tidak ada",
		})
	}

	var jawabanPeserta []models.JawabanPeserta

	// fmt.Println(data.IdUser)

	result := database.GormDB.Where("id_quiz = ? AND id_user = ?", data.IdQuiz, data.IdUser).Find(&jawabanPeserta)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// database.GormDB.Model(&jawabanPeserta).Association("Quiz").Find(&jawabanPeserta.Quiz)

	return c.Status(fiber.StatusOK).JSON(jawabanPeserta)
}
