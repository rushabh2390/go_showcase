package lead

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rushabh2390/go_showcase/go-fiber-crm/database"
	"gorm.io/gorm"
)

// Lead represents a lead in your application.
// @Description Lead model

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}
type InputLead struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
}

type DisplayLead struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Company   string     `json:"company"`
}

type CreateLeadResponse struct {
	Success bool        `json:"success"`
	Lead    DisplayLead `json:"lead",omitempty`
	Message string      `json:"message"`
}

// @Summary Get leads
// @Description Get all leads
// @Tags leads
// @Accept json
// @Produce json
// @Success 200 {array} DisplayLead
// @Router /api/v1/lead [get]
func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

// @Summary Get a lead
// @Description Get a lead by ID
// @Tags leads
// @Accept json
// @Produce json
// @Param id path int true "Lead ID"
// @Success 200 {object} DisplayLead
// @Router  /api/v1/lead/{id} [get]
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	if lead.Name == "" {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "No leads found with ID",
		})
		return nil
	}
	c.JSON(lead)
	return nil
}

// @Summary Create a new lead
// @Description Create a new lead
// @Tags leads
// @Accept json
// @Produce json
// @Param lead body InputLead true "New Lead"
// @Success 201 {object} CreateLeadResponse
// @Router /api/v1/lead [post]
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

// @Summary Delete a lead
// @Description Delete a lead by ID
// @Tags leads
// @Accept json
// @Produce json
// @Param id path int true "Lead ID"
// @Success 200 {object} CreateLeadResponse
// @Failure 400 {object} CreateLeadResponse
// @Router /api/v1/lead/{id} [delete]
func DeleteLeads(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "No leads found with ID",
		})
	}
	db.Delete(&lead)
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Lead Successfully Deleted",
	})
	return nil
}
