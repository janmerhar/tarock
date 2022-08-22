package board

import (
	"github.com/janmerhar/tarock/internal/deck"
	"github.com/janmerhar/tarock/internal/player"
	"github.com/janmerhar/tarock/internal/scoreboard"
)

type Board struct {
	Deck           deck.Deck       // deck of 54 tarock cards
	Cards_on_table deck.Deck       // currently placed cards on teble/board
	Talon          deck.Deck       // cards in talon
	Players        []player.Player // all the player playing current game
	Licitators     []player.Player // all the players that play for the points, it is a subset of Board.Players
	// Licitiotions: trula, kralji, itd.
	// Klic igre: 1, 2, ..., solo, klop,...

	scoreboard scoreboard.Scoreboard
}
