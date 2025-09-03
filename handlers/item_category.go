package handlers

import (
	"golink/models"
	"golink/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ItemCategory Handlers
func (handler *ItemHandler) CreateCategory(context *fiber.Ctx) error {
	var category models.ItemCategory
	if err := context.BodyParser(&category); err != nil {
		return utils.ErrorResponse(context, fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(category); err != nil {
		return utils.ErrorResponse(context, fiber.StatusBadRequest, err.Error())
	}

	if err := handler.repo.CreateCategory(&category); err != nil {
		return utils.ErrorResponse(context, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(context, fiber.StatusCreated, category)
}

func (handler *ItemHandler) GetCategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	category, err := handler.repo.GetCategoryByID(id)
	if err != nil {
		if err.Error() == "category not found" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    category,
	})
}

func (handler *ItemHandler) GetAllCategories(context *fiber.Ctx) error {
	categories, total, err := handler.repo.GetAllCategories()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    categories,
		"total":   total,
	})
}

func (handler *ItemHandler) UpdateCategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	var category models.ItemCategory
	if err := context.BodyParser(&category); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	category.ID = id
	if err := handler.repo.UpdateCategory(&category); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    category,
	})
}

func (handler *ItemHandler) DeleteCategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	if err := handler.repo.DeleteCategory(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
