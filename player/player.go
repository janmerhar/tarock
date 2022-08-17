package player

import (
	"github.com/janmerhar/tarock/cards"
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

// Put card on table
func (p *Player) Remove_from_hand(cardIndex int) cards.Card {
	return p.Hand.Draw_an_cards(cardIndex, 1)[0]
}

// Receive cards from talon/dealer
// by providing slice of cards to be added to player's hand
func (p *Player) Add_to_hand(new_cards []cards.Card) {
	p.Hand.Add_to_deck(new_cards)
}
