package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) makePretty(uglyMove *schemas.InternalMove) schemas.Move {
	color := uglyMove.Color
	piece := uglyMove.Piece
	from := uglyMove.From
	to := uglyMove.To
	flags := uglyMove.Flags
	captured := uglyMove.Captured
	promotion := uglyMove.Promotion

	var prettyFlags string

	for flag := range constants.Bits {
		if (constants.Bits[flag] & flags) != 0 {
			prettyFlags += constants.Flags[flag]
		}
	}

	fromAlgebraic := utils.Ox88ToSan(from)
	toAlgebraic := utils.Ox88ToSan(to)

	move := schemas.Move{
		Color:  color,
		Piece:  piece,
		From:   fromAlgebraic,
		To:     toAlgebraic,
		San:    c.moveToSan(uglyMove, c.moves(true, "", "")),
		Flags:  prettyFlags,
		Lan:    fromAlgebraic + toAlgebraic,
		Before: c.Fen(),
		After:  "",
	}

	// Generate the FEN for the 'after' key
	c.makeMove(uglyMove)
	move.After = c.Fen()
	c.undoMove()

	if captured != "" {
		move.Captured = captured
	}

	if promotion != "" {
		move.Promotion = promotion
		move.Lan += promotion
	}

	return move
}
