package utils

import (
	"testing"

	"github.com/mohitk09/cards_game/constants"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveCards(t *testing.T) {
	cards := RetrieveCards(constants.Numbers, constants.Suits)
	assert.Equal(t, len(cards), constants.Numbers*constants.Suits)
	assert.Equal(t, cards[0], int32(0))             // the first number in the array
	assert.Equal(t, cards[len(cards)-1], int32(51)) // the last number in the array

}

func TestConvertCodeToID(t *testing.T) {
	ID, err := ConvertCodeToID("AD")
	assert.Nil(t, err)
	assert.Equal(t, ID, int32(13))
}

func TestConvertCodeToIDWithInvalidSuit(t *testing.T) {
	_, err := ConvertCodeToID("A4")
	assert.Equal(t, "invalid suit", err.Error())
}

func TestConvertCodeToIDWithInvalidValue(t *testing.T) {
	_, err := ConvertCodeToID("ZS")
	assert.Equal(t, "invalid value", err.Error())
}

func TestRetrieveSelectedCards(t *testing.T) {
	IDs, err := RetrieveSelectedCards("AS,7H")
	assert.Nil(t, err)
	assert.Equal(t, len(IDs), 2)       // Only 2 cards in the deck
	assert.Equal(t, IDs[0], int32(0))  // For Ace of Spades
	assert.Equal(t, IDs[1], int32(45)) // For 7 of Hearts

}
