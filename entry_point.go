package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mohitk09/cards_game/api"
	"github.com/mohitk09/cards_game/database"
	"github.com/mohitk09/cards_game/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("cards_game.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	SetupRoutes(db)
}

func SetupRoutes(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&types.Deck{})

	app := fiber.New()

	app.Use(cors.New())

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server running")
	})

	deckRepository := database.NewDeckRepository(db)
	deckHandler := api.NewDeckHandler(deckRepository)

	app.Post("/deck", deckHandler.CreateDeck)
	app.Get("/deck/:id", deckHandler.OpenDeck)
	app.Get("/deck/:id/draw", deckHandler.Draw)

	log.Fatal(app.Listen(":8080"))
}
