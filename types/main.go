package types

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	ID         uuid.UUID
	IsShuffled bool
	Cards      pq.Int32Array `gorm:"type:integer[]"`
}

func (d *Deck) LeftOverCards() int {
	return len(d.Cards)
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	d.IsShuffled = true
}
