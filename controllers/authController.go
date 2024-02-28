package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nobitayon/QuizProject/database"
	middleware_jwt "github.com/nobitayon/QuizProject/middleware/middleware_jwt"
	"github.com/nobitayon/QuizProject/models"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func Register(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	// hashing with salt
	hashedSaltPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	createdUser := models.User{
		Nama:     data.Nama,
		Email:    data.Email,
		Password: string(hashedSaltPassword),
		Role:     "user",
	}

	result := database.GormDB.Create(&createdUser)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	claims := jwt.MapClaims{
		"id":   createdUser.Id,
		"role": createdUser.Role,
	}

	token, err := middleware_jwt.Encode(&claims, "secretKey")

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Nama:  createdUser.Nama,
		Email: createdUser.Email,
		Role:  createdUser.Role,
		Token: token,
	})
}

func RegisterAdmin(c *fiber.Ctx) error {

	claimsContext := c.Locals("jwtClaims").(jwt.MapClaims)

	if claimsContext["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "tidak diizinkan"})
	}

	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	// hashing with salt
	hashedSaltPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	createdUser := models.User{
		Nama:     data.Nama,
		Email:    data.Email,
		Password: string(hashedSaltPassword),
		Role:     "admin",
	}

	result := database.GormDB.Create(&createdUser)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	claims := jwt.MapClaims{
		"id":   createdUser.Id,
		"role": createdUser.Role,
	}

	token, err := middleware_jwt.Encode(&claims, "secretKey")

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Nama:  createdUser.Nama,
		Email: createdUser.Email,
		Role:  createdUser.Role,
		Token: token,
	})
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	var user models.User

	database.GormDB.Where("email = ?", data.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user tidak ditemukan",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {

		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password salah",
		})
	}

	claims := jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
	}

	token, err := middleware_jwt.Encode(&claims, "secretKey")

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Nama:  user.Nama,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	})
}
