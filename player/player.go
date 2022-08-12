package player

import (
	"github.com/janmerhar/tarock/deck"
)

type Player struct {
	Name string    // player's in game name
	Hand deck.Deck // player's cards that he holds in hand
}

// Constructor
func NewPlayer(name string) *Player {
	return &Player{Name: name, Hand: *deck.NewDeck()}
}
