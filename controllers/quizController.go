package controller

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/models"
)

type InputQuiz struct {
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
	WaktuMulai   string `json:"waktu_mulai"`
	WaktuSelesai string `json:"waktu_selesai"`
}

func AddQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data InputQuiz

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	waktuMulai, _ := time.Parse("2006-01-02 15:04:05", data.WaktuMulai)
	waktuSelesai, _ := time.Parse("2006-01-02 15:04:05", data.WaktuSelesai)

	createdQuiz := models.Quiz{
		Judul:        data.Judul,
		Deskripsi:    data.Deskripsi,
		WaktuMulai:   waktuMulai,
		WaktuSelesai: waktuSelesai,
	}

	result := database.GormDB.Create(&createdQuiz)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(createdQuiz)
}

func EditQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data models.Quiz

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada body request"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan penulisan path parameter id quiz"})
	}

	var quiz models.Quiz

	database.GormDB.Where("id = ?", idQuiz).First(&quiz)

	if quiz.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "quiz tidak ditemukan",
		})
	}

	result := database.GormDB.Model(&quiz).Updates(data)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(quiz)
}

func DeleteQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada penulisan path parameter id quiz"})
	}

	var quiz models.Quiz

	database.GormDB.Where("id = ?", idQuiz).First(&quiz)

	if quiz.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "tidak ditemukan quiz",
		})
	}

	result := database.GormDB.Delete(&quiz)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(quiz)
}

func GetQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	idQuiz, err := strconv.Atoi(c.Params("idQuiz"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "kesalahan pada path parameter id quiz"})
	}

	var quiz models.Quiz

	database.GormDB.Where("id = ?", idQuiz).First(&quiz)

	if quiz.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "quiz tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(quiz)
}

func GetAllQuiz(c *fiber.Ctx) error {

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var quizs []models.Quiz

	result := database.GormDB.Find(&quizs)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(quizs)
}
