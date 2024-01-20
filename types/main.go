package types

import (
	"math/rand"

	"github.com/lib/pq"
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

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}
