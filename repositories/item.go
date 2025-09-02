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

func (repo *ItemRepository) CreateItem(item *models.Item) error {
	return repo.db.Create(item).Error
}

func (repo *ItemRepository) GetItemByID(id uint64) (*models.Item, error) {
	var item models.Item
	err := repo.db.Preload("Subcategory.Category").Preload("ItemImages").First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return &item, nil
}

func (repo *ItemRepository) GetAllItems(limit int, offset int, sort string, sort_direction string) ([]models.Item, int64, error) {
	var items []models.Item
	var total int64

	if err := repo.db.Model(&models.Item{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := repo.db.Preload("Subcategory.Category").Preload("ItemImages").Order(sort + " " + sort_direction).Limit(limit).Offset(offset).Find(&items).Error
	return items, total, err
}

func (repo *ItemRepository) UpdateItem(id uint64, item *map[string]interface{}) error {
	return repo.db.Model(&models.Item{}).Where("id = ?", id).Updates(item).Error
}

func (repo *ItemRepository) DeleteItem(id uint64) error {
	return repo.db.Delete(&models.Item{}, id).Error
}

func (repo *ItemRepository) CreateItemImage(image *models.ItemImages) error {
	return repo.db.Create(image).Error
}

func (repo *ItemRepository) GetItemImages(itemID uint64) ([]models.ItemImages, error) {
	var images []models.ItemImages
	err := repo.db.Where("item_id = ?", itemID).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
func (repo *ItemRepository) GetItemImagesByID(imageID uint64) (*models.ItemImages, error) {
	var image models.ItemImages
	err := repo.db.First(&image, imageID).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}
func (repo *ItemRepository) DeleteItemImage(imageID uint64) error {
	image, err := repo.GetItemImagesByID(imageID)
	if err == nil && image != nil {
		os.Remove("./assets/images/items/" + image.Image)
	}
	return repo.db.Where("id = ?", imageID).Delete(&models.ItemImages{}).Error
}

// ItemSubcategory CRUD Operations

func (repo *ItemRepository) CreateSubcategory(subcategory *models.ItemSubcategory) error {
	return repo.db.Create(subcategory).Error
}

func (repo *ItemRepository) GetSubcategoryByID(id uint64) (*models.ItemSubcategory, error) {
	var subcategory models.ItemSubcategory
	err := repo.db.Preload("Category").First(&subcategory, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subcategory not found")
		}
		return nil, err
	}
	return &subcategory, nil
}

func (repo *ItemRepository) GetSubcategoryByCategoryID(id uint64) ([]models.ItemSubcategory, error) {
	var subcategory []models.ItemSubcategory
	err := repo.db.Preload("Category").Where("category_id = ?", id).Find(&subcategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subcategory not found")
		}
		return nil, err
	}
	return subcategory, nil
}

func (repo *ItemRepository) GetAllSubcategories() ([]models.ItemSubcategory, error) {
	var subcategories []models.ItemSubcategory
	err := repo.db.Preload("Category").Find(&subcategories).Error
	return subcategories, err
}

func (repo *ItemRepository) UpdateSubcategory(subcategory *models.ItemSubcategory) error {
	return repo.db.Save(subcategory).Error
}

func (repo *ItemRepository) DeleteSubcategory(id uint64) error {
	return repo.db.Delete(&models.ItemSubcategory{}, id).Error
}

// ItemCategory CRUD Operations

func (repo *ItemRepository) CreateCategory(category *models.ItemCategory) error {
	return repo.db.Create(category).Error
}

func (repo *ItemRepository) GetCategoryByID(id uint64) (*models.ItemCategory, error) {
	var category models.ItemCategory
	err := repo.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (repo *ItemRepository) GetAllCategories() ([]models.ItemCategory, error) {
	var categories []models.ItemCategory
	err := repo.db.Find(&categories).Error
	return categories, err
}

func (repo *ItemRepository) UpdateCategory(category *models.ItemCategory) error {
	return repo.db.Save(category).Error
}

func (repo *ItemRepository) DeleteCategory(id uint64) error {
	return repo.db.Delete(&models.ItemCategory{}, id).Error
}
