package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/database"
	_ "github.com/rushabh2390/go_showcase/go-fiber-crm/docs"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead/:id", lead.DeleteLeads)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("conenction opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	log.Fatal(app.Listen(":8000"))
	// defer database.DBConn.Close()
}
