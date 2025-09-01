package repositories

import (
	"errors"
	"golink/models"
	"os"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

// Item CRUD Operations

func (r *ItemRepository) CreateItem(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepository) GetItemByID(id uint64) (*models.Item, error) {
	var item models.Item
	err := r.db.Preload("Subcategory.Category").Preload("ItemImages").First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetAllItems(limit int, offset int, sort string, sort_direction string) ([]models.Item, int64, error) {
	var items []models.Item
	var total int64

	if err := r.db.Model(&models.Item{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Subcategory.Category").Preload("ItemImages").Order(sort + " " + sort_direction).Limit(limit).Offset(offset).Find(&items).Error
	return items, total, err
}

func (r *ItemRepository) UpdateItem(id uint64, item *map[string]interface{}) error {
	return r.db.Model(&models.Item{}).Where("id = ?", id).Updates(item).Error
}

func (r *ItemRepository) DeleteItem(id uint64) error {
	return r.db.Delete(&models.Item{}, id).Error
}

func (r *ItemRepository) CreateItemImage(image *models.ItemImages) error {
	return r.db.Create(image).Error
}

func (r *ItemRepository) GetItemImages(itemID uint64) ([]models.ItemImages, error) {
	var images []models.ItemImages
	err := r.db.Where("item_id = ?", itemID).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
func (r *ItemRepository) GetItemImagesByID(imageID uint64) (*models.ItemImages, error) {
	var image models.ItemImages
	err := r.db.First(&image, imageID).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}
func (r *ItemRepository) DeleteItemImage(imageID uint64) error {
	image, err := r.GetItemImagesByID(imageID)
	if err == nil && image != nil {
		os.Remove("./assets/images/items/" + image.Image)
	}
	return r.db.Where("id = ?", imageID).Delete(&models.ItemImages{}).Error
}

// ItemSubcategory CRUD Operations

func (r *ItemRepository) CreateSubcategory(subcategory *models.ItemSubcategory) error {
	return r.db.Create(subcategory).Error
}

func (r *ItemRepository) GetSubcategoryByID(id uint64) (*models.ItemSubcategory, error) {
	var subcategory models.ItemSubcategory
	err := r.db.Preload("Category").First(&subcategory, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subcategory not found")
		}
		return nil, err
	}
	return &subcategory, nil
}

func (r *ItemRepository) GetSubcategoryByCategoryID(id uint64) ([]models.ItemSubcategory, error) {
	var subcategory []models.ItemSubcategory
	err := r.db.Preload("Category").Where("category_id = ?", id).Find(&subcategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subcategory not found")
		}
		return nil, err
	}
	return subcategory, nil
}

func (r *ItemRepository) GetAllSubcategories() ([]models.ItemSubcategory, error) {
	var subcategories []models.ItemSubcategory
	err := r.db.Preload("Category").Find(&subcategories).Error
	return subcategories, err
}

func (r *ItemRepository) UpdateSubcategory(subcategory *models.ItemSubcategory) error {
	return r.db.Save(subcategory).Error
}

func (r *ItemRepository) DeleteSubcategory(id uint64) error {
	return r.db.Delete(&models.ItemSubcategory{}, id).Error
}

// ItemCategory CRUD Operations

func (r *ItemRepository) CreateCategory(category *models.ItemCategory) error {
	return r.db.Create(category).Error
}

func (r *ItemRepository) GetCategoryByID(id uint64) (*models.ItemCategory, error) {
	var category models.ItemCategory
	err := r.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *ItemRepository) GetAllCategories() ([]models.ItemCategory, error) {
	var categories []models.ItemCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *ItemRepository) UpdateCategory(category *models.ItemCategory) error {
	return r.db.Save(category).Error
}

func (r *ItemRepository) DeleteCategory(id uint64) error {
	return r.db.Delete(&models.ItemCategory{}, id).Error
}
