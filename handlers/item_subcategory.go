package handlers

import (
	"golink/models"
	"golink/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ItemSubcategory Handlers

func (handler *ItemHandler) CreateSubcategory(context *fiber.Ctx) error {
	var subcategory models.ItemSubcategory
	if err := context.BodyParser(&subcategory); err != nil {
		return utils.ErrorResponse(context, fiber.StatusBadRequest, err.Error())
	}

	if err := utils.ValidateStruct(context, &subcategory); err != nil {
		return utils.ErrorResponse(context, fiber.StatusBadRequest, err.Error())
	}

	if err := handler.repo.CreateSubcategory(&subcategory); err != nil {
		return utils.ErrorResponse(context, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(context, fiber.StatusCreated, subcategory)
}

func (handler *ItemHandler) GetSubcategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	subcategory, err := handler.repo.GetSubcategoryByID(id)
	if err != nil {
		if err.Error() == "subcategory not found" {
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
		"data":    subcategory,
	})
}

func (handler *ItemHandler) GetSubcategoryByCategoryID(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	subcategory, err := handler.repo.GetSubcategoryByCategoryID(id)
	if err != nil {
		if err.Error() == "subcategory not found" {
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
		"data":    subcategory,
	})
}

func (handler *ItemHandler) GetAllSubcategories(context *fiber.Ctx) error {
	subcategories, err := handler.repo.GetAllSubcategories()
	if err != nil {
		return utils.ErrorResponse(context, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(context, fiber.StatusOK, subcategories)
}

func (handler *ItemHandler) UpdateSubcategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	var subcategory models.ItemSubcategory
	if err := context.BodyParser(&subcategory); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	subcategory.ID = id
	if err := handler.repo.UpdateSubcategory(&subcategory); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    subcategory,
	})
}

func (handler *ItemHandler) DeleteSubcategory(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	if err := handler.repo.DeleteSubcategory(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"message": "Subcategory deleted successfully",
	})
}
