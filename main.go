package main

import (
	"fmt"

	"contacts-api-sqlite/contacts"
	"contacts-api-sqlite/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func routes(app *fiber.App) {
	app.Get("/api/v1/contacts", contacts.GetContacts)
	app.Get("/api/v1/contact/:id", contacts.GetContact)
	app.Post("/api/v1/contact", contacts.NewContact)
	app.Delete("/api/v1/contact/:id", contacts.DeleteContact)
	app.Put("/api/v1/contact/:id", contacts.UpdateContact)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&contacts.Contact{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	// defer database.DBConn.Close()
	routes(app)
	app.Listen(":8000")
}
