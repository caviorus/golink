package handlers

import (
	"golink/config"
	"golink/models"
	"golink/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(context *fiber.Ctx) error {
	var input models.User
	if err := context.BodyParser(&input); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid credentials",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Token error",
		})
	}

	return context.JSON(fiber.Map{
		"status": true,
		"token":  tokenString,
	})
}
