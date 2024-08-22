package chess

import (
	"fmt"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) Fen() string {
	var empty int
	var fen string

	for i := constants.Ox88["a8"]; i <= constants.Ox88["h1"]; i++ {
		if c.board[i] != nil {
			if empty > 0 {
				fen += fmt.Sprint(empty)
				empty = 0
			}

			color := c.board[i].Color
			pieceType := c.board[i].PieceType

			if color == constants.ColorWhite {
				fen += strings.ToUpper(pieceType)
			} else {
				fen += strings.ToLower(pieceType)
			}
		} else {
			empty++
		}

		if ((i + 1) & 0x88) != 0 {
			if empty > 0 {
				fen += fmt.Sprint(empty)
			}

			if i != constants.Ox88["h1"] {
				fen += "/"
			}

			empty = 0
			i += 8
		}
	}

	var castling string

	if (c.castling[constants.ColorWhite] & constants.Bits[constants.BitKingSideCastle]) != 0 {
		castling += "K"
	}

	if (c.castling[constants.ColorWhite] & constants.Bits[constants.BitQueenSideCastle]) != 0 {
		castling += "Q"
	}

	if (c.castling[constants.ColorBlack] & constants.Bits[constants.BitKingSideCastle]) != 0 {
		castling += "k"
	}

	if (c.castling[constants.ColorBlack] & constants.Bits[constants.BitQueenSideCastle]) != 0 {
		castling += "q"
	}

	// Do we have an empty castling flag?
	if castling == "" {
		castling = "-"
	}

	epSquare := "-"

	// Only print the ep square if en passant is a valid move (Pawn is present and ep capture is not pinned)
	if c.epSquare != constants.EmptySquare {
		bigPawnSquare := c.epSquare

		if c.turn == constants.ColorWhite {
			bigPawnSquare += 16
		} else {
			bigPawnSquare -= 16
		}

		squares := []int{bigPawnSquare + 1, bigPawnSquare - 1}

		for _, square := range squares {
			// Is the square off the board?
			if (square & 0x88) != 0 {
				continue
			}

			color := c.turn

			// Is there a pawn that can capture the en passant square?
			if c.board[square] != nil && c.board[square].Color == color &&
				c.board[square].PieceType == constants.PiecePawn {
				// If the pawn makes an ep capture, does it leave it's king in check?
				c.makeMove(&schemas.InternalMove{
					Color:    color,
					From:     square,
					To:       c.epSquare,
					Piece:    constants.PiecePawn,
					Captured: constants.PiecePawn,
					Flags:    constants.Bits[constants.BitEpCapture],
				})

				isLegal := !c.isKingAttacked(color)
				c.undoMove()

				// If ep is legal, break and set the ep square in the FEN output
				if isLegal {
					epSquare = utils.Ox88ToSan(c.epSquare)
					break
				}
			}
		}
	}

	return strings.Join([]string{fen, c.turn, castling, epSquare, fmt.Sprint(c.halfMoves), fmt.Sprint(c.moveNumber)}, " ")
}
