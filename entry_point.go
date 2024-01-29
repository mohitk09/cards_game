package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
)

// func main() {

// 	db, err := gorm.Open(sqlite.Open("cards_game.db"), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	app := SetupRoutes(db)
// 	log.Fatal(app.Listen(":8080"))
// }

// func SetupRoutes(db *gorm.DB) *fiber.App {
// 	// Migrate the schema
// 	db.AutoMigrate(&types.Deck{})

// 	app := fiber.New()

// 	app.Use(cors.New())

// 	// Health check
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Server running")
// 	})

// 	// Setup the DB layer
// 	deckRepository := database.NewDeckRepository(db)

// 	// Setup the endpoint layer
// 	deckHandler := api.NewDeckHandler(deckRepository)

// 	app.Post("/deck", deckHandler.CreateDeck)
// 	app.Get("/deck/:id", deckHandler.OpenDeck)
// 	app.Get("/deck/:id/draw", deckHandler.Draw)

// 	return app

// }

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hola")
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
