package api

import (
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

	var cardIds []int32
	if cardsQuery == "" {
		cardIds = utils.RetrieveCards(constants.Numbers, constants.Suits)
	} else {
		var err error
		cardIds, err = utils.RetrieveSelectedCards(cardsQuery)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
	}

	deck := new(types.Deck)
	deck.ID = uuid.New().String() // convert to string and store for the ease
	deck.Cards = cardIds

	if shuffle {
		deck.Shuffle()
	}

	_, err := handler.repository.Create(*deck)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "failed creating the deck",
			"error":   err,
		})
	}

	return c.Status(http.StatusCreated).JSON(deck.CreateDeckResponse())
}

func (handler *DeckHandler) OpenDeck(c *fiber.Ctx) error {
	return c.JSON("Something is right")
}

func (handler *DeckHandler) Draw(c *fiber.Ctx) error {
	return c.JSON("hell yeah")
}
