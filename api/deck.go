package api

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mohitk09/cards_game/constants"
	"github.com/mohitk09/cards_game/database"
	"github.com/mohitk09/cards_game/types"
	"github.com/mohitk09/cards_game/utils"
)

/* API handler */
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
	deck, err := handler.repository.Find(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(deck.OpenDeckResponse())
}

func (handler *DeckHandler) Draw(c *fiber.Ctx) error {
	count, err := strconv.Atoi(c.Query("count", "1"))

	if err != nil || count < 1 {
		return c.JSON(http.StatusBadRequest, "count param invalid, please pass a value greater than or equal to 1")
	}

	deck, err := handler.repository.Find(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	remainingCards := deck.LeftOverCards()

	// Send 409 in case the user tries to draw from an empty deck
	if remainingCards == 0 {
		return c.Status(http.StatusConflict).JSON("empty deck, please draw from a different deck")
	}

	servePartial := false

	if count > remainingCards {
		count = remainingCards
		remainingCards = 0
		servePartial = true
	}

	var cards pq.Int32Array
	// Iterate from the back of the slice as drawing a card simulates a stack operation
	for count > 0 && len(deck.Cards) >= 0 {
		cards = append(cards, deck.Cards[len(deck.Cards)-1])
		deck.Cards = deck.Cards[:len(deck.Cards)-1]
		count--
	}

	// Update the database
	handler.repository.Save(deck)

	// Sends 206 as the request can't be fully served
	if servePartial {
		return c.Status(http.StatusPartialContent).JSON(deck.DrawCardResponse(cards))

	}

	return c.Status(http.StatusOK).JSON(deck.DrawCardResponse(cards))
}
