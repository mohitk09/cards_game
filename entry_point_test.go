package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetDeck(t *testing.T) {
	tests := []struct {
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			route:        "/deck/something_random",
			expectedCode: 404,
		},
	}

	db, _ := gorm.Open(sqlite.Open("cards_game_test.db"), &gorm.Config{})

	app := SetupRoutes(db)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}
}
