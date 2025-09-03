package handlers

import (
	"fmt"
	"golink/models"
	"golink/repositories"
	"golink/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	repo *repositories.ItemRepository
}

func NewItemHandler(repo *repositories.ItemRepository) *ItemHandler {
	return &ItemHandler{repo: repo}
}

// Item Handlers

func (itemHandler *ItemHandler) CreateItem(context *fiber.Ctx) error {
	var data MapInterface
	if err := context.BodyParser(&data); err != nil {
		return context.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	name := data["name"].(string)
	subcategoryID := data["subcategory_id"].(float64)
	link := data["link"].(string)
	description := data["description"].(string)
	_type := data["type"].(string)

	subcategoryIDUint, err := strconv.ParseUint(fmt.Sprintf("%.0f", subcategoryID), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid subcategory_id",
		})
	}

	item := models.Item{
		Name:          name,
		SubcategoryID: subcategoryIDUint,
		Link:          link,
		Description:   description,
		Type:          _type,
	}

	if err := itemHandler.repo.CreateItem(&item); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    item,
	})
}

func (itemHandler *ItemHandler) GetItem(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	item, err := itemHandler.repo.GetItemByID(id)
	if err != nil {
		if err.Error() == "item not found" {
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

	item.Thumbnail = utils.CreateImageLink(item.Thumbnail)
	for i := 0; i < len(item.ItemImages); i++ {
		item.ItemImages[i].Image = utils.CreateImageLink(item.ItemImages[i].Image)
	}
	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    item,
	})
}

func (itemHandler *ItemHandler) GetAllItems(context *fiber.Ctx) error {
	limit, err := strconv.Atoi(context.Query("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(context.Query("start", "0"))
	if err != nil {
		offset = 0
	}
	sort := context.Query("sort", "name")
	sort_direction := context.Query("sortDirection", "asc")

	items, total, err := itemHandler.repo.GetAllItems(limit, offset, sort, sort_direction)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    items,
		"total":   total,
	})
}

func (itemHandler *ItemHandler) UpdateItem(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	name := context.FormValue("name")
	link := context.FormValue("link")
	_type := context.FormValue("type")
	description := context.FormValue("description")
	subcategoryIDStr := context.FormValue("subcategory_id")
	subcategoryIDUint, err := strconv.ParseUint(subcategoryIDStr, 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid subcategory_id",
		})
	}

	if err := itemHandler.repo.UpdateItem(id, &map[string]interface{}{
		"name":           name,
		"subcategory_id": subcategoryIDUint,
		"link":           link,
		"description":    description,
		"type":           _type,
	}); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
	})
}

func (itemHandler *ItemHandler) DeleteItem(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	if err := itemHandler.repo.DeleteItem(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Item deleted successfully",
	})
}
