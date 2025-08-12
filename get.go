package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) Get(square string) *schemas.Piece {
	pieceOnBoard := c.board[constants.Ox88[square]]

	return pieceOnBoard
}
