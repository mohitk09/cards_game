package types

import (
	"math/rand"
	"strconv"

	"github.com/lib/pq"
	"github.com/mohitk09/cards_game/constants"
	"gorm.io/gorm"
)

type CreateDeckResponse struct {
	DeckId     string `json:"deck_id"` // returns string representation of UUID
	IsShuffled bool   `json:"shuffled"`
	Remaining  int32  `json:"remaining"`
}

type Deck struct {
	gorm.Model
	ID         string // Convert UUID to string and store
	IsShuffled bool
	Cards      pq.Int32Array `gorm:"type:integer[]"`
}

func (d *Deck) LeftOverCards() int {
	return len(d.Cards)
}

func (d *Deck) CreateDeckResponse() CreateDeckResponse {
	return CreateDeckResponse{d.ID, d.IsShuffled, int32(d.LeftOverCards())}
}

type CardResponse struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type DrawCardResponse struct {
	Cards []CardResponse `json:"cards"`
}

func (d *Deck) DrawCardResponse(cards pq.Int32Array) DrawCardResponse {
	return DrawCardResponse{Cards: convertIDsToCardJson(cards)}
}

type OpenDeckResponse struct {
	CreateDeckResponse
	DrawCardResponse
}

func (d *Deck) OpenDeckResponse() OpenDeckResponse {
	return OpenDeckResponse{d.CreateDeckResponse(), d.DrawCardResponse(d.Cards)}
}

func convertIDsToCardJson(ids []int32) (cardResponse []CardResponse) {
	for _, id := range ids {
		cardResponse = append(cardResponse, ConvertIDToCode(id))
	}
	return
}

func ConvertIDToCode(id int32) (cardResponse CardResponse) {
	// Get the value
	value := id % (constants.Numbers)
	switch value {
	case 0:
		cardResponse.Value = "ACE"
	case 10:
		cardResponse.Value = "JACK"
	case 11:
		cardResponse.Value = "QUEEN"
	case 12:
		cardResponse.Value = "KING"
	default:
		cardResponse.Value = strconv.Itoa(int(value) + 1)
	}
	cardResponse.Code = cardResponse.Value[:1]

	/* The following order suit order is maintained to compute the Suit
	Spades, Diamonds, Clubs and Hearts
	*/
	switch id / (constants.Numbers) {
	case 0:
		cardResponse.Suit = "SPADES"
	case 1:
		cardResponse.Suit = "DIAMONDS"
	case 2:
		cardResponse.Suit = "CLUBS"
	case 3:
		cardResponse.Suit = "HEARTS"
	}
	cardResponse.Code += cardResponse.Suit[:1]

	return
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}
