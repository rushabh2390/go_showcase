package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/database"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead:id", lead.DeleteLeads)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
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
	app.Listen("8000")
	defer database.DBConn.Close()
}
