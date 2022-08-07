package player

import "github.com/janmerhar/tarock/cards"

type Player struct {
	Name              string       // player's in game name
	Hand              []cards.Card // player's cards that he holds in hand
}
