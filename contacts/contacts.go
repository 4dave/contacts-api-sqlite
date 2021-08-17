package contacts

import (
	"contacts-api/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

func GetContacts(c *fiber.Ctx) error {
	db := database.DBConn
	var contact []Contact
	db.Find(&contact)
	return c.JSON(contact)
}

func GetContact(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var contact Contact
	db.Find(&contact, id)
	return c.JSON(contact)
}

func NewContact(c *fiber.Ctx) error {
	db := database.DBConn
	contact := new(Contact)
	if err := c.BodyParser(contact); err != nil {
		c.SendStatus(500)
	}
	db.Create(&contact)
	return c.JSON(contact)
}

func EditContact(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	contact := new(Contact)
	if err := c.BodyParser(contact); err != nil {
		c.SendStatus(500)
	}
	db.First(&contact, id)
	db.Save(&contact)
	return c.JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var contact Contact
	db.First(&contact, id)
	if contact.ID == 0 {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": "no record found with given ID"})
		return nil
	}
	db.Delete(&contact)
	return c.JSON(fiber.Map{"success": "record successfully deleted"})
}
