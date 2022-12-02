package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"
)

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
	defer func() {
		sqlDB, err := database.DBConn.DB()
		if err != nil {
			log.Fatalln(err)
		}

		sqlDB.Close()
	}()
}

func setupRoutes(app *fiber.App) {
	app.Get("/leads", lead.GetLeads)

	app.Get("/leads/:id", lead.GetLead)

	app.Post("/leads", lead.NewLead)

	app.Delete("/leads/:id", lead.DeleteLead)
}

func initDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connection opened to db")
	db.AutoMigrate(&lead.Lead{})
	database.DBConn = db
	log.Println("DB Migrated")
}
