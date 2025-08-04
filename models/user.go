package models

type User struct {
	BaseID
	Name     string
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
}
