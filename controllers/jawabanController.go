package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nobitayon/QuizProject/database"
	"github.com/nobitayon/QuizProject/models"
)

func AddJawaban(c *fiber.Ctx) error {

	var jawaban models.JawabanPeserta

	if err := c.BodyParser(&jawaban); err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	claims := c.Locals("jwtClaims").(jwt.MapClaims)

	if claims["id"] != jawaban.IdUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	result := database.GormDB.Create(&jawaban)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jawaban)
}

func EditJawaban(c *fiber.Ctx) error {

	var data models.JawabanPeserta

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"message": "error parsing"})
	}

	var jawaban models.JawabanPeserta

	database.GormDB.Where("id = ?", id).First(&jawaban)

	if jawaban.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "jawaban tidak ditemukan",
		})
	}

	result := database.GormDB.Model(&jawaban).Updates(data)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jawaban)
}

func GetJawabanById(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"message": "error parsing"})
	}

	var jawaban models.JawabanPeserta

	database.GormDB.Where("id = ?", id).First(&jawaban)

	if jawaban.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "jawaban tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(jawaban)
}

func GetJawaban(c *fiber.Ctx) error {

	var jawabans []models.JawabanPeserta

	result := database.GormDB.Find(&jawabans)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jawabans)
}

func DeleteJawabanById(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"message": "error parsing"})
	}

	var jawaban models.JawabanPeserta

	database.GormDB.Where("id = ?", id).First(&jawaban)

	if jawaban.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "jawaban tidak ditemukan",
		})
	}

	result := database.GormDB.Delete(&jawaban)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jawaban)
}
