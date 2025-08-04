package seed

import (
	"golink/models"
	"golink/utils"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	password, err := utils.HashPassword("password")
	if err != nil {
		return err
	}
	users := []models.User{
		{
			Name:     "Caviorus",
			Username: "caviorus",
			Email:    "xlucief@gmail.com",
			Phone:    "6289652660352",
			Password: password,
		},
	}

	for _, r := range users {
		var existing models.User
		err := db.Where("username = ?", r.Username).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if c_err := db.Create(&r).Error; c_err != nil {
				return c_err
			}
		}

		if err == nil {
			existing.Name = r.Name
			existing.Email = r.Email
			existing.Phone = r.Phone
			existing.Password = r.Password
			if err := db.Save(&existing).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
