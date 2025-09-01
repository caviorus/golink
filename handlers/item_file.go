package handlers

import (
	"fmt"
	"golink/models"
	"golink/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// func to upload items images
func (h *ItemHandler) UploadItemImages(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
			"error":   err.Error(),
		})
	}

	item, err := h.repo.GetItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "item not found",
		})
	}

	name := item.Name

	thumbnail, err := c.FormFile("thumbnail")
	if err == nil {
		ext := filepath.Ext(thumbnail.Filename)
		thumbnailName := fmt.Sprintf("%s-%s%s",
			utils.GenerateSlug(name),
			utils.RandomString(8),
			ext,
		)
		path := fmt.Sprintf("./assets/images/items/%s", thumbnailName)
		_ = c.SaveFile(thumbnail, path)
		h.repo.UpdateItem(id, &map[string]interface{}{
			"thumbnail": thumbnailName,
		})
	}

	form, err := c.MultipartForm()
	if err == nil && form != nil {
		images := form.File["images[]"]
		for _, image := range images {
			ext := filepath.Ext(image.Filename)
			imageName := fmt.Sprintf("%s-%s%s",
				utils.GenerateSlug(name),
				utils.RandomString(8),
				ext,
			)
			path := fmt.Sprintf("./assets/images/items/%s", imageName)
			_ = c.SaveFile(image, path)

			//save file name to itemImage
			h.repo.CreateItemImage(&models.ItemImages{
				ItemID: item.ID,
				Image:  imageName,
			})
		}
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Images uploaded successfully",
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

	image, err := h.repo.GetItemImagesByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	os.Remove(image.Image)
	h.repo.DeleteItemImage(image.ID)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Image deleted successfully",
	})
}
