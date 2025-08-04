package models

type Item struct {
	BaseID
	Name          string          `gorm:"type:varchar(100)" json:"name"`
	Type          string          `gorm:"type:varchar(12)" json:"type"`
	Link          string          `gorm:"type:varchar(100)" json:"link"`
	Description   string          `gorm:"type:text" json:"description"`
	Thumbnail     string          `gorm:"type:varchar(255)" json:"thumbnail"`
	SubcategoryID uint64          `json:"subcategory_id" gorm:"not null"`
	Subcategory   ItemSubcategory `gorm:"foreignKey:SubcategoryID" json:"subcategory,omitempty"`
}

type ItemSubcategory struct {
	BaseID
	Name       string       `gorm:"type:varchar(100)" json:"name"`
	CategoryID uint64       `json:"category_id" gorm:"not null"`
	Category   ItemCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

type ItemCategory struct {
	BaseID
	Name string `gorm:"type:varchar(100)" json:"name"`
}
