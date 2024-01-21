package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohitk09/cards_game/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test create deck, a map to reduce redundancy
func TestCreateDeck(t *testing.T) {
	tests := []struct {
		route         string
		expectedCode  int // expected HTTP status code
		numberOfCards int // expected number of cards in the deck
	}{
		{
			route:         "/deck",
			expectedCode:  http.StatusCreated,
			numberOfCards: 52,
		},
		{
			route:         "/deck?cards=AS,7H",
			expectedCode:  http.StatusCreated,
			numberOfCards: 2,
		},
	}

	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})
	app := SetupRoutes(db)
	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, nil)
		res, err := app.Test(req, 100000)

		assert.Nil(t, err)
		assert.Equal(t, test.expectedCode, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		var jsonResponse map[string]json.RawMessage
		json.Unmarshal(body, &jsonResponse)
		assert.Contains(t, jsonResponse, "deck_id")
		assert.Contains(t, jsonResponse, "shuffled")
		assert.Contains(t, jsonResponse, "remaining")

		var response types.CreateDeckResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			assert.Equal(t, false, response.IsShuffled)
			assert.Equal(t, test.numberOfCards, response.Remaining)

		}
	}
}

func TestCreateDeckMalformedCase(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})
	app := SetupRoutes(db)
	req := httptest.NewRequest("POST", "/deck?cards=YZ", nil)
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

}

func TestOpenDeckRecordNotFound(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})
	app := SetupRoutes(db)
	req := httptest.NewRequest("GET", "/deck/something_random", nil)
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

}
