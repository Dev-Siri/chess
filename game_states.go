package chess

import "github.com/Dev-Siri/chess/constants"

func (c *Chess) IsCheck() bool {
	return c.isKingAttacked(c.turn)
}

func (c *Chess) IsCheckmate() bool {
	availableMoves := len(c.moves(true, "", ""))
	return c.IsCheck() && availableMoves == 0
}

func (c *Chess) IsStalemate() bool {
	availableMoves := len(c.moves(true, "", ""))
	return !c.IsCheck() && availableMoves == 0
}

func (c *Chess) IsInsufficientMaterial() bool {
	pieces := map[string]int{
		constants.PieceBishop: 0,
		constants.PieceKnight: 0,
		constants.PieceRook:   0,
		constants.PieceQueen:  0,
		constants.PieceKing:   0,
		constants.PiecePawn:   0,
	}

	bishops := []int{}
	numPieces, squareColor := 0, 0

	for i := constants.Ox88["a8"]; i <= constants.Ox88["h1"]; i++ {
		squareColor = (squareColor + 1) % 2

		if (i & 0x88) != 0 {
			i += 7
			continue
		}

		piece := c.board[i]

		if piece != nil {
			if _, exists := pieces[piece.PieceType]; exists {
				pieces[piece.PieceType] = pieces[piece.PieceType] + 1
			} else {
				pieces[piece.PieceType] = 1
			}

			if piece.PieceType == constants.PieceBishop {
				bishops = append(bishops, squareColor)
			}

			numPieces++
		}
	}

	// K vs. K
	if numPieces == 2 {
		return true
	} else if numPieces == 3 &&
		(pieces[constants.PieceBishop] == 1 || pieces[constants.PieceKnight] == 1) {
		// K vs. KN or K vs. KB
		return true
	} else if numPieces == (pieces[constants.PieceBishop] + 2) {
		// KB vs. KB where any number of bishops are all on the same color
		var sum int

		bishopsLen := len(bishops)

		for i := 0; i < bishopsLen; i++ {
			sum += bishops[i]
		}

		if sum == 0 || sum == bishopsLen {
			return true
		}
	}

	return false
}

func (c *Chess) IsThreeFoldRepetition() bool {
	return c.getPositionCount(c.Fen()) >= 3
}

func (c *Chess) IsDraw() bool {
	return c.halfMoves >= 100 || c.IsStalemate() || c.IsInsufficientMaterial() || c.IsThreeFoldRepetition()
}

func (c *Chess) IsGameOver() bool {
	return c.IsCheckmate() || c.IsStalemate() || c.IsDraw()
}
