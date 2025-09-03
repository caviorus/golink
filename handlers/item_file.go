package handlers

import (
	"fmt"
	"golink/models"
	"golink/utils"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// func to upload items images
func (handler *ItemHandler) UploadItemImages(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.FormValue("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
			"error":   err.Error(),
		})
	}

	item, err := handler.repo.GetItemByID(id)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "item not found",
		})
	}

	name := item.Name

	thumbnail, err := context.FormFile("thumbnail")
	if err == nil {
		ext := filepath.Ext(thumbnail.Filename)
		thumbnailName := fmt.Sprintf("%s-%s%s",
			utils.GenerateSlug(name),
			utils.RandomString(8),
			ext,
		)
		path := fmt.Sprintf("./assets/images/items/%s", thumbnailName)
		_ = context.SaveFile(thumbnail, path)
		handler.repo.UpdateItem(id, &map[string]interface{}{
			"thumbnail": thumbnailName,
		})
	}

	form, err := context.MultipartForm()
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
			_ = context.SaveFile(image, path)

			//save file name to itemImage
			handler.repo.CreateItemImage(&models.ItemImages{
				ItemID: item.ID,
				Image:  imageName,
			})
		}
	}

	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Images uploaded successfully",
	})

}

func (handler *ItemHandler) DeleteItemImage(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	image, err := handler.repo.GetItemImagesByID(id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	handler.repo.DeleteItemImage(image.ID)
	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Image deleted successfully",
	})
}

func (handler *ItemHandler) DeleteItemThumbnail(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "invalid ID",
		})
	}

	item, err := handler.repo.GetItemByID(id)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	handler.repo.DeleteItemThumbnail(item.ID)
	return context.JSON(fiber.Map{
		"status":  true,
		"message": "Image deleted successfully",
	})
}
