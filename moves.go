package chess

import (
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) moves(legal bool, piece string, square string) []schemas.InternalMove {
	var forSquare string
	forPiece := strings.ToLower(piece)

	if square != "" {
		forSquare = strings.ToLower(square)
	}

	moves := []schemas.InternalMove{}
	us := c.turn
	them := utils.SwapColors(us)

	firstSquare := constants.Ox88["a8"]
	lastSquare := constants.Ox88["h1"]
	singleSquare := false

	// Are we generating moves for a single square?
	if forSquare != "" {
		// Illegal square, return empty moves
		if _, exists := constants.Ox88[forSquare]; !exists {
			return []schemas.InternalMove{}
		} else {
			firstSquare = constants.Ox88[forSquare]
			singleSquare = true
		}
	}

	for from := firstSquare; from <= lastSquare; from++ {
		// Did we run off the end of the board
		if (from & 0x88) != 0 {
			from += 7
			continue
		}

		// Empty square or opponent, skip
		if c.board[from] == nil || c.board[from].Color == them {
			continue
		}

		pieceType := c.board[from].PieceType

		var to int

		if pieceType == constants.PiecePawn {
			if forPiece != "" && forPiece != pieceType {
				continue
			}

			// Single square, non-capturing
			to = from + constants.PawnOffsets[us][0]
			if c.board[to] == nil {
				utils.AddMove(&moves, us, from, to, pieceType, "", 0)

				// Double square
				to = from + constants.PawnOffsets[us][1]
				if constants.SecondRank[us] == utils.Rank(from) && c.board[to] == nil {
					utils.AddMove(&moves, us, from, to, constants.PiecePawn, "", constants.Bits[constants.BitBigPawn])
				}
			}

			// Pawn captures
			for j := 2; j < 4; j++ {
				to = from + constants.PawnOffsets[us][j]

				if (to & 0x88) != 0 {
					continue
				}

				if c.board[to] != nil && c.board[to].Color == them {
					utils.AddMove(&moves, us, from, to, constants.PiecePawn, c.board[to].PieceType, constants.Bits[constants.BitCapture])
				} else if to == c.epSquare {
					utils.AddMove(&moves, us, from, to, constants.PiecePawn, constants.PiecePawn, constants.Bits[constants.BitCapture])
				}
			}
		} else {
			if forPiece != "" && forPiece != pieceType {
				continue
			}

			pieceOffsetLen := len(constants.PieceOffsets[pieceType])

			for j := 0; j < pieceOffsetLen; j++ {
				offset := constants.PieceOffsets[pieceType][j]
				to = from

				for {
					to += offset

					if (to & 0x88) != 0 {
						break
					}

					if c.board[to] == nil {
						utils.AddMove(&moves, us, from, to, pieceType, "", 0)
					} else {
						if c.board[to].Color == us {
							break
						}

						utils.AddMove(&moves, us, from, to, pieceType, c.board[to].PieceType, constants.Bits[constants.BitCapture])
						break
					}

					if pieceType == constants.PieceKnight || pieceType == constants.PieceKing {
						break
					}
				}
			}
		}
	}

	// Check for castling if we're generating all moves,
	// or doing single square move generation on the king's square.
	if forPiece == "" || forPiece == constants.PieceKing {
		if !singleSquare || lastSquare == c.kings[us] {
			// King-Side castling
			if (c.castling[us] & constants.Bits[constants.BitKingSideCastle]) != 0 {
				castlingFrom := c.kings[us]
				castlingTo := castlingFrom + 2

				kingSquareAttacked, _ := c.attacked(them, c.kings[us])
				castlingFromSquareAttacked, _ := c.attacked(them, castlingFrom+1)
				castlingToSquareAttacked, _ := c.attacked(them, castlingTo)

				if c.board[castlingFrom+1] == nil &&
					c.board[castlingTo] == nil &&
					!kingSquareAttacked &&
					!castlingFromSquareAttacked &&
					!castlingToSquareAttacked {
					utils.AddMove(&moves, us, c.kings[us], castlingTo, constants.PieceKing, "", constants.Bits[constants.BitKingSideCastle])
				}
			}

			if (c.castling[us] & constants.Bits[constants.BitQueenSideCastle]) != 0 {
				castlingFrom := c.kings[us]
				castlingTo := castlingFrom - 2

				kingSquareAttacked, _ := c.attacked(them, c.kings[us])
				castlingFromSquareAttacked, _ := c.attacked(them, castlingFrom-1)
				castlingToSquareAttacked, _ := c.attacked(them, castlingTo)

				if c.board[castlingFrom-1] == nil &&
					c.board[castlingFrom-2] == nil &&
					c.board[castlingFrom-3] == nil &&
					!kingSquareAttacked &&
					!castlingFromSquareAttacked &&
					!castlingToSquareAttacked {
					utils.AddMove(&moves, us, c.kings[us], castlingTo, constants.PieceKing, "", constants.Bits[constants.BitQueenSideCastle])
				}
			}
		}
	}

	if !legal || c.kings[us] == constants.EmptySquare {
		return moves
	}

	legalMoves := []schemas.InternalMove{}

	movesLen := len(moves)

	for i := 0; i < movesLen; i++ {
		c.makeMove(&moves[i])

		if !c.isKingAttacked(us) {
			legalMoves = append(legalMoves, moves[i])
		}

		c.undoMove()
	}

	return legalMoves
}

func (c *Chess) Moves(square string, piece string) ([]string, []schemas.Move) {
	moves := c.moves(true, piece, square)
	prettyMoves := []schemas.Move{}
	sanMoves := []string{}

	for _, move := range moves {
		prettyMoves = append(prettyMoves, c.makePretty(&move))
		sanMoves = append(sanMoves, c.moveToSan(&move, moves))
	}

	return sanMoves, prettyMoves
}
