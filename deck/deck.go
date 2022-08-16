package deck

import (
	"math/rand"
	"sort"
	"time"

	"github.com/janmerhar/tarock/cards"
)

type Deck struct {
	Cards []cards.Card // current cards in deck of 54
}

// Constructors for Deck
// https://stackoverflow.com/questions/18125625/constructors-in-go
func NewDeck() *Deck {
	return &Deck{append([]cards.Card{}, cards.AllCards...)}
}

func NewCustomDeck(new_cards []cards.Card) *Deck {
	return &Deck{append([]cards.Card{}, new_cards...)}
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

// Sort cards in order by colors: tarock, heart, spade, diamond, clover
// and then by power
func (deck *Deck) Sort_cards() {
	// Order of sorting elements by color
	priority_by_color := map[string]int{
		"tarock":  5,
		"heart":   4,
		"spade":   3,
		"diamond": 2,
		"clover":  1,
	}

	sort.Slice(deck.Cards, func(a, b int) bool {
		// Sort cards by color
		if deck.Cards[a].Type_of_card != deck.Cards[b].Type_of_card {
			return priority_by_color[deck.Cards[a].Type_of_card] > priority_by_color[deck.Cards[b].Type_of_card]
		}
		// and then by power
		return deck.Cards[a].Power > deck.Cards[b].Power
	})

}

// Add cards to deck at index
func (deck *Deck) Add_to_deck(new_cards []cards.Card) {
	deck.Cards = append(deck.Cards, new_cards...)
}
