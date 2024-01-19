package database

import (
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
