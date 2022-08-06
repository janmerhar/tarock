package deck

import (
	"math/rand"
	"time"

	"github.com/janmerhar/tarock/cards"
)

type Deck struct {
	Cards []cards.Card // current cards in deck of 54
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
}

func (deck *Deck) Draw_n_cards(n int) []cards.Card {
	subslice := deck.Cards[:n]
	deck.Cards = deck.Cards[n:]

	return subslice
}

// draw cards from range to range
// koristno, ko izbiras iz talona
