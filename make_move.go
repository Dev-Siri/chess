package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) makeMove(move *schemas.InternalMove) {
	us := c.turn
	them := utils.SwapColors(us)

	c.pushToHistory(move)

	c.board[move.To] = c.board[move.From]
	c.board[move.From] = nil

	// If en-passant capture, remove the captured pawn
	if (move.Flags & constants.Bits[constants.BitEpCapture]) != 0 {
		if c.turn == constants.ColorBlack {
			c.board[move.To-16] = nil
		} else {
			c.board[move.To+16] = nil
		}
	}

	// If pawn promotion, replace with new piece
	if move.Promotion != "" {
		c.board[move.To] = &schemas.Piece{
			PieceType: move.Promotion,
			Color:     us,
		}
	}

	// If we moved the king
	if c.board[move.To].PieceType == constants.PieceKing {
		c.kings[us] = move.To

		// If we castled, move the rook next to the king
		if (move.Flags & constants.Bits[constants.BitKingSideCastle]) != 0 {
			castlingTo := move.To - 1
			castlingFrom := move.To + 1

			c.board[castlingTo] = c.board[castlingFrom]
			c.board[castlingFrom] = nil
		} else if (move.Flags & constants.Bits[constants.BitQueenSideCastle]) != 0 {
			castlingTo := move.To + 1
			castlingFrom := move.To - 2

			c.board[castlingTo] = c.board[castlingFrom]
			c.board[castlingFrom] = nil
		}

		// Turn off castling
		c.castling[us] = 0
	}

	// Turn off castling if we move a rook
	if c.castling[us] != 0 {
		for i := 0; i < len(constants.Rooks[us]); i++ {
			if move.From == constants.Rooks[us][i]["square"] && ((c.castling[us] & constants.Rooks[us][i]["flag"]) != 0) {
				c.castling[us] ^= constants.Rooks[us][i]["flag"]
				break
			}
		}
	}

	// Turn off castlign if we capture a rook
	if c.castling[them] != 0 {
		for i := 0; i < len(constants.Rooks[them]); i++ {
			if move.To == constants.Rooks[them][i]["square"] && ((c.castling[them] & constants.Rooks[them][i]["flag"]) != 0) {
				c.castling[them] ^= constants.Rooks[them][i]["flag"]
				break
			}
		}
	}

	// If big pawn move, update the en passant square
	if (move.Flags & constants.Bits[constants.BitBigPawn]) != 0 {
		if us == constants.ColorBlack {
			c.epSquare = move.To - 16
		} else {
			c.epSquare = move.To + 16
		}
	} else {
		c.epSquare = constants.EmptySquare
	}

	// Reset the 50 move counter if a pawn is moved or a piece is captured
	if move.Piece == constants.PiecePawn {
		c.halfMoves = 0
	} else if (move.Flags & (constants.Bits[constants.BitCapture] | constants.Bits[constants.BitEpCapture])) != 0 {
		c.halfMoves = 0
	} else {
		c.halfMoves++
	}

	if us == constants.ColorBlack {
		c.moveNumber++
	}

	c.turn = them
}
