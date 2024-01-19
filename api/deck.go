package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohitk09/cards_game/database"
)

/* API handler, which is called by an instance of the DB object */

type DeckHandler struct {
	repository *database.DeckRepository
}

func NewDeckHandler(repository *database.DeckRepository) *DeckHandler {
	return &DeckHandler{
		repository: repository,
	}
}

func (handler *DeckHandler) CreateDeck(c *fiber.Ctx) error {
	return c.JSON("hello there")
}

func (handler *DeckHandler) OpenDeck(c *fiber.Ctx) error {
	return c.JSON("Something is right")
}

func (handler *DeckHandler) Draw(c *fiber.Ctx) error {
	return c.JSON("hell yeah")
}
