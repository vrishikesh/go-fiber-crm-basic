package lead

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	database.DBConn.Find(&leads)
	c.Status(http.StatusOK).JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.Find(&lead, id)
	c.Status(http.StatusOK).JSON(lead)
	return nil
}

func NewLead(c *fiber.Ctx) error {
	lead := &Lead{}
	if err := c.BodyParser(&lead); err != nil {
		c.Status(http.StatusInternalServerError).JSON(err)
		return err
	}
	database.DBConn.Create(&lead)
	c.Status(http.StatusCreated).JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.Delete(&lead, id)
	c.Status(http.StatusNoContent)
	return nil
}
