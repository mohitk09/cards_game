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

// Test create deck, a map to reduce repeating several lines of code
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
		res, err := app.Test(req)

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

func TestOpenDeckRecord(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})
	app := SetupRoutes(db)

	req := httptest.NewRequest("POST", "/deck", nil)
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	var createResponse types.CreateDeckResponse
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	json.Unmarshal(body, &createResponse)

	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId, nil)
	openDeckRes, _ := app.Test(req)
	openDeckResBody, openDeckBodyErr := ioutil.ReadAll(openDeckRes.Body)
	assert.Nil(t, openDeckBodyErr)
	var openDeckResponse types.OpenDeckResponse
	json.Unmarshal(openDeckResBody, &openDeckResponse)
	assert.Equal(t, http.StatusOK, openDeckRes.StatusCode)
}

func TestDrawDeck(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})
	app := SetupRoutes(db)

	req := httptest.NewRequest("POST", "/deck", nil)
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	var createResponse types.CreateDeckResponse
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	json.Unmarshal(body, &createResponse)

	req, _ = http.NewRequest("GET", "/deck/"+createResponse.DeckId+"/draw", nil)
	drawDeckRes, _ := app.Test(req)
	openDeckResBody, drawDeckBodyErr := ioutil.ReadAll(drawDeckRes.Body)
	assert.Nil(t, drawDeckBodyErr)
	var drawDeckResponse types.DrawCardResponse
	json.Unmarshal(openDeckResBody, &drawDeckResponse)
	assert.Equal(t, http.StatusOK, drawDeckRes.StatusCode)
}
