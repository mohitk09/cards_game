package database

import (
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
