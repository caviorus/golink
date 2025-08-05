package repositories

import (
	"errors"
	"golink/models"

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
	err := r.db.Preload("Subcategory.Category").First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetAllItems(limit int, offset int) ([]models.Item, int64, error) {
	var items []models.Item
	var total int64

	if err := r.db.Model(&models.Item{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Subcategory.Category").Limit(limit).Offset(offset).Find(&items).Error
	return items, total, err
}

func (r *ItemRepository) UpdateItem(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *ItemRepository) DeleteItem(id uint64) error {
	return r.db.Delete(&models.Item{}, id).Error
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
