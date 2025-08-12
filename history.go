package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) pushToHistory(move *schemas.InternalMove) {
	c.history = append(c.history, schemas.History{
		Move: move,
		Kings: map[string]int{
			constants.ColorWhite: c.kings[constants.ColorWhite],
			constants.ColorBlack: c.kings[constants.ColorBlack],
		},
		Turn: c.turn,
		Castling: map[string]int{
			constants.ColorWhite: c.castling[constants.ColorWhite],
			constants.ColorBlack: c.castling[constants.ColorBlack],
		},
		EpSquare:   c.epSquare,
		HalfMoves:  c.halfMoves,
		MoveNumber: c.moveNumber,
	})
}
