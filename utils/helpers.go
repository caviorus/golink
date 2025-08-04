package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ArrayColumn(data []map[string]interface{}, key string) []string {
	var result []string
	for _, item := range data {
		if val, ok := item[key].(string); ok {
			result = append(result, val)
		}
	}
	return result
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  false,
		"message": message,
		"data":    nil,
	})
}

func SuccessResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}

var validate = validator.New()

func ValidateStruct(c *fiber.Ctx, s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return nil
}
