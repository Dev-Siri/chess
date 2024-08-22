package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) undoMove() *schemas.InternalMove {
	newHistory, old := utils.PopMove(c.history)

	c.history = newHistory

	if old.Turn == "" {
		return nil
	}

	move := old.Move

	c.kings = old.Kings
	c.turn = old.Turn
	c.castling = old.Castling
	c.epSquare = old.EpSquare
	c.halfMoves = old.HalfMoves
	c.moveNumber = old.MoveNumber

	us := c.turn
	them := utils.SwapColors(us)

	c.board[move.From] = c.board[move.To]
	c.board[move.From].PieceType = move.Piece // To undo any promotions

	c.board[move.To] = nil

	if move.Captured != "" {
		if (move.Flags & constants.Bits[constants.BitEpCapture]) != 0 {
			// En passant capture
			var index int

			if us == constants.ColorBlack {
				index = move.To - 16
			} else {
				index = move.To + 16
			}

			c.board[index] = &schemas.Piece{
				Color:     them,
				PieceType: constants.PiecePawn,
			}
		} else {
			// Regular capture
			c.board[move.To] = &schemas.Piece{
				PieceType: move.Captured,
				Color:     them,
			}
		}
	}

	if (move.Flags & (constants.Bits[constants.BitKingSideCastle] | constants.Bits[constants.BitQueenSideCastle])) != 0 {
		var castlingTo, castlingFrom int

		if (move.Flags & constants.Bits[constants.BitKingSideCastle]) != 0 {
			castlingTo = move.To + 1
			castlingFrom = move.To - 1
		} else {
			castlingTo = move.To - 2
			castlingFrom = move.To + 1
		}

		c.board[castlingTo] = c.board[castlingFrom]
		c.board[castlingFrom] = nil
	}

	return move
}

func (c *Chess) Undo() *schemas.Move {
	move := c.undoMove()

	if move != nil {
		prettyMove := c.makePretty(move)
		c.decPositionCount(prettyMove.After)
		return &prettyMove
	}

	return nil
}
