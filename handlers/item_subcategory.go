package handlers

import (
	"golink/models"
	"golink/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ItemSubcategory Handlers

func (h *ItemHandler) CreateSubcategory(c *fiber.Ctx) error {
	var subcategory models.ItemSubcategory
	if err := c.BodyParser(&subcategory); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := utils.ValidateStruct(c, &subcategory); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.repo.CreateSubcategory(&subcategory); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, subcategory)
}

func (h *ItemHandler) GetSubcategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	subcategory, err := h.repo.GetSubcategoryByID(id)
	if err != nil {
		if err.Error() == "subcategory not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    subcategory,
	})
}

func (h *ItemHandler) GetSubcategoryByCategoryID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	subcategory, err := h.repo.GetSubcategoryByCategoryID(id)
	if err != nil {
		if err.Error() == "subcategory not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    subcategory,
	})
}

func (h *ItemHandler) GetAllSubcategories(c *fiber.Ctx) error {
	subcategories, err := h.repo.GetAllSubcategories()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, subcategories)
}

func (h *ItemHandler) UpdateSubcategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	var subcategory models.ItemSubcategory
	if err := c.BodyParser(&subcategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	subcategory.ID = id
	if err := h.repo.UpdateSubcategory(&subcategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    subcategory,
	})
}

func (h *ItemHandler) DeleteSubcategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "invalid ID",
		})
	}

	if err := h.repo.DeleteSubcategory(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Subcategory deleted successfully",
	})
}
