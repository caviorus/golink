package handlers

import (
	"fmt"
	"golink/models"
	"golink/repositories"
	"golink/utils"
	"os"
	"path/filepath"
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
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if err := utils.ValidateStruct(c, &item); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
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

	items, total, err := h.repo.GetAllItems(limit, offset)
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

	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	item.ID = id
	if err := h.repo.UpdateItem(&item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"input":   item,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    item,
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

// func to upload items images
func (h *ItemHandler) UploadItemImage(c *fiber.Ctx) error {
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

	// Ambil form multipart
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data",
		})
	}

	// Ambil semua file dari field "images[]"
	files := form.File["images"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded under 'images[]'",
		})
	}

	// Buat folder upload jika belum ada
	uploadDir := "./assets/images/items"
	os.MkdirAll(uploadDir, os.ModePerm)

	// Simpan semua file dan simpan info untuk response
	var uploadedFiles []fiber.Map
	itemName := utils.GenerateSlug(item.Name)

	for _, file := range files {
		dst := filepath.Join(uploadDir, itemName, file.Filename)
		if err := c.SaveFile(file, dst); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to save file %s", file.Filename),
			})
		}
		uploadedFiles = append(uploadedFiles, fiber.Map{
			"filename": file.Filename,
			"path":     dst,
			"size":     file.Size,
		})
		h.repo.CreateItemImage(&models.ItemImages{
			ItemID: id,
			Type:   "PATH",
			URL:    dst,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Files uploaded successfully",
		"files":   uploadedFiles,
	})
}

func (h *ItemHandler) GetItemImages(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	images, err := h.repo.GetItemImages(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    images,
	})
}

func (h *ItemHandler) DeleteItemImage(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	_imageID := c.Params("image_id")
	if _imageID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "image ID is required",
		})
	}

	imageID, err := strconv.ParseUint(_imageID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid image ID",
		})
	}

	image, err := h.repo.GetItemImagesByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if image.ID == imageID {
		os.Remove(image.URL)
		h.repo.DeleteItemImage(image.ID)
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "Image deleted successfully",
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  false,
		"message": "Image not found",
	})
}
