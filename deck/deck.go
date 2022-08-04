package deck

import (
	"math/rand"
	"time"

	"github.com/janmerhar/tarock/cards"
)

type Deck struct {
	Cards []cards.Card // current cards in deck of 54
}

func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
}
