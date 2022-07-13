package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/subratkumarsahu/go-fiber-crm-basic/database"
	"github.com/subratkumarsahu/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect the database")
	}

	fmt.Println("connection opened to database")
	// database.DBconn.AutoMigrate(&lead.Lead{})
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}
