package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

func NewLeads(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&lead)
	c.Status(201).JSON(&fiber.Map{
		"success": true,
		"lead":    lead,
		"message": "Lead created successfuly",
	})
	return nil
}

func DeleteLeads(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   "No leads found with ID",
		})
	}
	db.Delete(&lead)
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Lead Successfully Deleted",
	})
	return nil
}
