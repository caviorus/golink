package routes

import (
	"golink/handlers"
	"golink/middleware"
	"golink/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/login", handlers.Login)

	// Initialize repositories and handlers
	itemRepo := repositories.NewItemRepository(db)
	itemHandler := handlers.NewItemHandler(itemRepo)

	api := app.Group("/api", middleware.Access())

	// Item routes
	items := api.Group("/items")
	items.Post("/", itemHandler.CreateItem)
	items.Get("/", itemHandler.GetAllItems)
	items.Get("/:id", itemHandler.GetItem)
	items.Put("/:id", itemHandler.UpdateItem)
	items.Delete("/:id", itemHandler.DeleteItem)

	// Subcategory routes
	subcategories := api.Group("/subcategories")
	subcategories.Post("/", itemHandler.CreateSubcategory)
	subcategories.Get("/", itemHandler.GetAllSubcategories)
	subcategories.Get("/:id", itemHandler.GetSubcategory)
	subcategories.Put("/:id", itemHandler.UpdateSubcategory)
	subcategories.Delete("/:id", itemHandler.DeleteSubcategory)

	// Category routes
	categories := api.Group("/categories")
	categories.Post("/", itemHandler.CreateCategory)
	categories.Get("/", itemHandler.GetAllCategories)
	categories.Get("/:id", itemHandler.GetCategory)
	categories.Put("/:id", itemHandler.UpdateCategory)
	categories.Delete("/:id", itemHandler.DeleteCategory)
}
