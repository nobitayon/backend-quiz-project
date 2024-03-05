package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/models"
)

type InputPertanyaan struct {
	Id           int    `json:"id"`
	Pertanyaan   string `json:"pertanyaan"`
	OpsiJawaban  string `json:"opsiJawaban"`
	JawabanBenar int    `json:"jawabanBenar"`
	IdQuiz       int    `json:"idQuiz"`
}

func EditPertanyaanOnQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data models.Pertanyaan

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan penulisan id quiz"})
	}

	idPertanyaan, err := strconv.Atoi(c.Params("idPertanyaan"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan penulisan id pertanyaan"})
	}

	var pertanyaan models.Pertanyaan

	database.GormDB.Where("id_quiz = ? AND id = ?", idQuiz, idPertanyaan).First(&pertanyaan)

	if pertanyaan.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "pertanyaan yang ingin diedit tidak ditemukan",
		})
	}

	result := database.GormDB.Model(&pertanyaan).Updates(data)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	database.GormDB.Model(&pertanyaan).Association("Quiz").Find(&pertanyaan.Quiz)

	return c.Status(fiber.StatusOK).JSON(pertanyaan)
}

func CreatePertanyaanOnQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data models.Pertanyaan

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan penulisan id quiz"})
	}

	var quiz models.Quiz

	database.GormDB.Where("id = ?", idQuiz).First(&quiz)

	if quiz.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "gagal menambah pertanyaan, karena quiz tempat pertanyaan akan ditambahkan tidak ada",
		})
	}

	createdPertanyaan := models.Pertanyaan{
		Pertanyaan:   data.Pertanyaan,
		OpsiJawaban:  data.OpsiJawaban,
		JawabanBenar: data.JawabanBenar,
		IdQuiz:       idQuiz,
	}

	result := database.GormDB.Create(&createdPertanyaan).Preload("Quiz")

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	database.GormDB.Model(&createdPertanyaan).Association("Quiz").Find(&createdPertanyaan.Quiz)

	return c.Status(fiber.StatusOK).JSON(createdPertanyaan)
}

func DeletePertanyaanOnQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data models.Pertanyaan

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan penulisan id quiz"})
	}

	idPertanyaan, err := strconv.Atoi(c.Params("idPertanyaan"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada penulisan id pertanyaan"})
	}

	var pertanyaan models.Pertanyaan

	database.GormDB.Where("id_quiz = ? AND id = ?", idQuiz, idPertanyaan).First(&pertanyaan)

	if pertanyaan.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "gagal menghapus pertanyaan, karena pertanyaan tidak ditemukan",
		})
	}

	result := database.GormDB.Delete(&pertanyaan)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	database.GormDB.Model(&pertanyaan).Association("Quiz").Find(&pertanyaan.Quiz)

	return c.Status(fiber.StatusOK).JSON(pertanyaan)
}

func GetPertanyaanOnQuiz(c *fiber.Ctx) error {

	var data models.Pertanyaan

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada penulisan id quiz"})
	}

	idPertanyaan, err := strconv.Atoi(c.Params("idPertanyaan"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada penulisan id pertanyaan"})
	}

	var pertanyaan models.Pertanyaan

	database.GormDB.Where("id_quiz = ? AND id = ?", idQuiz, idPertanyaan).Preload("Quiz").First(&pertanyaan)

	if pertanyaan.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "tidak terdapat pertanyaan yang diminta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(pertanyaan)
}

func GetAllPertanyaanOnQuiz(c *fiber.Ctx) error {

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	var pertanyaans []models.Pertanyaan

	result := database.GormDB.Where("id_quiz = ?", idQuiz).Preload("Quiz").Find(&pertanyaans)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(pertanyaans)
}
