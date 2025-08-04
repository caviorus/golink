package main

import (
	"golink/config"
	"golink/models"
	"golink/routes"
	"golink/seed"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	db := config.DB

	// AutoMigrate models
	err := db.AutoMigrate(
		&models.User{},
		&models.ItemCategory{},
		&models.ItemSubcategory{},
		&models.Item{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	if err := seed.SeedUser(db); err != nil {
		log.Fatalf("failed seeding: %v", err)
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)
	log.Fatal(app.Listen(":8080"))
}
