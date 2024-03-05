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

func PeriksaQuizFromUser(c *fiber.Ctx) error {
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

	result := database.GormDB.Where("id_quiz = ? AND id_user = ?", data.IdQuiz, data.IdUser).Preload("Pertanyaan").Find(&jawabanPeserta)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	num_correct := 0

	var pertanyaans *[]models.Pertanyaan

	database.GormDB.Where("id_quiz = ?", data.IdQuiz).Find(&pertanyaans)
	num_pertanyaan := len(*pertanyaans)

	for _, jawab := range jawabanPeserta {
		if jawab.JawabanPeserta == jawab.Pertanyaan.JawabanBenar {
			data := models.JawabanPeserta{Skor: 1}
			database.GormDB.Model(&jawab).Updates(data)
			num_correct += 1
		}

	}

	nilai := (float64(num_correct) / float64(num_pertanyaan)) * 100

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"nilai": nilai, "jawabanPeserta": jawabanPeserta})
}
