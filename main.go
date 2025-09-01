package main

import (
	"golink/config"
	"golink/models"
	"golink/routes"
	"golink/seed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		&models.ItemImages{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	if err := seed.SeedUser(db); err != nil {
		log.Fatalf("failed seeding: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Ubah sesuai kebutuhan
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.SetupRoutes(app, db)
	log.Fatal(app.Listen(":8080"))
}
