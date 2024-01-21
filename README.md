# Cards Game

This project create Cards game APIs which could be used in Poker and Blackjack or any other card game.

## APIs

There are three APIs that have been written, which are scalable and could be used in any game.

```
GET  /                    # Health check
POST /deck                # This creates a new deck (optional params: cards, shuffle) -- cards
                            can also be passed as query param, the API would then just
                            create a deck with those cards
GET  /deck/:id            # Retrieves a deck by id, all the cards should be returned
GET  /deck/:id/draw       # Draw cards from the deck (optional params: count) -- The count is
                            the number of cards that should be drawn, default value is 1.
                            It simulates a stack operation
```

## Dependencies

1. [Fiber](https://docs.gofiber.io/) for setting up the routes
2. [Gorm](https://pkg.go.dev/gorm.io/gorm) for the ORM layer
3. [Sqllite3](https://pkg.go.dev/github.com/mattn/go-sqlite3) as choice of the DB

## Repository Structure

The Project is divided into several packages, **important ones are**:-

1. **Database** :- Has basic CRUD operations, uses Sqllite under the hood
2. **API** :- All three APIs reside in this folder
3. **Utils**:- Functions which are used across different files
4. **Types**:- All the types/structs and functions which are dependent on those types
5. There are also some test cases which have been added

## Run it locally

- Clone the repo, install the go dependencies. You could either do `go mod download` or `go get ./...`.
- Execute `go run entry_point.go` which starts the server by configuring DB and the routes.
- Check if http://localhost:8080 shows the message `Server running`. If yes the project setup is done.
- You can start by creating some decks and copying the UUIDs to later open the deck or draw cards.

## Future scope

1. More intensive test cases.
2. Creating a docker file so that it becomes easier to test and deploy.
