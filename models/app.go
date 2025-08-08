package models

type BaseID struct {
	ID *uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
}
