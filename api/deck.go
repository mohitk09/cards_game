package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mohitk09/cards_game/constants"
	"github.com/mohitk09/cards_game/database"
	"github.com/mohitk09/cards_game/types"
	"github.com/mohitk09/cards_game/utils"
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
	shuffleQuery := c.Query("shuffle")
	cardsQuery := c.Query("cards")

	shuffle := shuffleQuery == "true"

	var listOfCards []int32
	if cardsQuery == "" {
		listOfCards = utils.RetrieveCards(constants.Numbers, constants.Suits)
	}

	deck := new(types.Deck)
	deck.ID = uuid.New()
	deck.Cards = listOfCards

	if shuffle {
		deck.Shuffle()
	}

	ok, err := handler.repository.Create(*deck)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Failed creating item",
			"error":   err,
		})
	}

	fmt.Println("I want it to stop here", ok)
	return c.JSON(ok)
}

func (handler *DeckHandler) OpenDeck(c *fiber.Ctx) error {
	return c.JSON("Something is right")
}

func (handler *DeckHandler) Draw(c *fiber.Ctx) error {
	return c.JSON("hell yeah")
}
