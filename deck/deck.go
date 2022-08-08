package deck

import (
	"math/rand"
	"time"

	"github.com/janmerhar/tarock/cards"
)

type Deck struct {
	Cards []cards.Card // current cards in deck of 54
}

// Randomly shuffle contents of []deck.Deck
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
}

// Draw n cards from index [0] forwards
func (deck *Deck) Draw_n_cards(n int) []cards.Card {
	subslice := deck.Cards[:n]
	deck.Cards = deck.Cards[n:]

	return subslice
}

// Draw n cards from index a
func (deck *Deck) Draw_an_cards(a int, n int) []cards.Card {
	subslice := make([]cards.Card, n)

	// Copying cards into subslice
	for i := 0; i < n; i++ {
		subslice[i] = deck.Cards[a+i]
	}

	// Removing drawn cards from deck
	deck.Cards = append(deck.Cards[:a], deck.Cards[a+n:]...)
	return subslice
}
