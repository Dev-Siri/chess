package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) Clear(preserveHeaders bool) {
	c.board = [128]*schemas.Piece{}
	c.turn = constants.ColorWhite
	c.kings = map[string]int{
		constants.ColorWhite: constants.EmptySquare,
		constants.ColorBlack: constants.EmptySquare,
	}
	c.epSquare = constants.EmptySquare
	c.halfMoves = 0
	c.moveNumber = 1
	c.history = []schemas.History{}
	c.comments = make(map[string]string)
	c.positionCount = make(map[string]int)

	if !preserveHeaders {
		c.header = make(map[string]string)
	}

	// Delete the "SetUp" and "FEN" headers (if preserved), the board is empty and these headers
	// don't make sense in this state. They'll get added later via c.Load() or .put()
	delete(c.header, "SetUp")
	delete(c.header, "FEN")
}
