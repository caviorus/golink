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

func (h *ItemHandler) CreateItem(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	name := data["name"].(string)
	subcategoryID := data["subcategory_id"].(float64)
	link := data["link"].(string)
	description := data["description"].(string)
	_type := data["type"].(string)

	subcategoryIDUint, err := strconv.ParseUint(fmt.Sprintf("%.0f", subcategoryID), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	if err := h.repo.CreateItem(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    item,
	})
}

func (h *ItemHandler) GetItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	item, err := h.repo.GetItemByID(id)
	if err != nil {
		if err.Error() == "item not found" {
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

	item.Thumbnail = utils.CreateImageLink(item.Thumbnail)
	for i := 0; i < len(item.ItemImages); i++ {
		item.ItemImages[i].Image = utils.CreateImageLink(item.ItemImages[i].Image)
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    item,
	})
}

func (h *ItemHandler) GetAllItems(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.Query("start", "0"))
	if err != nil {
		offset = 0
	}
	sort := c.Query("sort", "name")
	sort_direction := c.Query("sortDirection", "asc")

	items, total, err := h.repo.GetAllItems(limit, offset, sort, sort_direction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data": fiber.Map{
			"items": items,
			"total": total,
		},
	})
}

func (h *ItemHandler) UpdateItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	name := c.FormValue("name")
	link := c.FormValue("link")
	_type := c.FormValue("type")
	description := c.FormValue("description")
	subcategoryIDStr := c.FormValue("subcategory_id")
	subcategoryIDUint, err := strconv.ParseUint(subcategoryIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid subcategory_id",
		})
	}

	if err := h.repo.UpdateItem(id, &map[string]interface{}{
		"name":           name,
		"subcategory_id": subcategoryIDUint,
		"link":           link,
		"description":    description,
		"type":           _type,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
	})
}

func (h *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	if err := h.repo.DeleteItem(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Item deleted successfully",
	})
}
