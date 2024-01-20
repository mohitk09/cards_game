package database

import (
	"errors"

	"github.com/mohitk09/cards_game/types"
	"gorm.io/gorm"
)

// This class deals with the basic CRUD operations i.e. DB layer of the application
type DeckRepository struct {
	db *gorm.DB
}

func NewDeckRepository(db *gorm.DB) *DeckRepository {
	return &DeckRepository{
		db: db,
	}
}

func (repository *DeckRepository) Create(deck types.Deck) (types.Deck, error) {
	err := repository.db.Create(&deck).Error
	if err != nil {
		return deck, err
	}

	return deck, nil
}

func (repository *DeckRepository) Find(ID string) (types.Deck, error) {
	var deck types.Deck
	deck.ID = ID
	err := repository.db.First(&deck).Error
	if err != nil {
		err = errors.New("deck not found")
	}
	return deck, err
}
