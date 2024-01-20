# Cards Game

This project create Cards game APIs which could be used in Poker and Blackjack or any other card game.

# APIs

There are three APIs that have written, which are scalable and could be used in any game.

```
GET  /                    # Health check
POST /deck                # This creates a new deck (optional params: cards, shuffle) -- cards can also be passed as query param, the API would then just create a deck with those cards
GET  /deck/:id            # Retrieves a deck by id
GET  /deck/:id/draw       # Draw cards from the deck (optional params: count) -- The count is the number of cards that should be drawn, default value is 1
```

# Dependencies

1. Fiber for setting up the routes
2. Gorm for the ORM layer
3. SQL Lite as choice of the DB

# Structure

The Project is divided into several packages, mainly:-

1. API
2. Database
3. Utils
4. Types
5. Constants
