package chess

import (
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) put(pieceType string, color string, square string) bool {
	if !strings.Contains(constants.PieceSymbols, strings.ToLower(pieceType)) {
		return false
	}

	if _, exists := constants.Ox88[square]; !exists {
		return false
	}

	sq := constants.Ox88[square]

	if pieceType == constants.PieceKing &&
		!(c.kings[color] == constants.EmptySquare || c.kings[color] == sq) {
		return false
	}

	currentPieceOnSquare := c.board[sq]

	if currentPieceOnSquare != nil && currentPieceOnSquare.PieceType == constants.PieceKing {
		c.kings[currentPieceOnSquare.Color] = constants.EmptySquare
	}

	c.board[sq] = &schemas.Piece{
		PieceType: pieceType,
		Color:     color,
	}

	if pieceType == constants.PieceKing {
		c.kings[color] = sq
	}

	return true
}

func (c *Chess) Put(pieceType string, color string, square string) bool {
	if c.put(pieceType, color, square) {
		c.updateCastlingRights()
		c.updateEnPassantSquare()
		c.updateSetup(c.Fen())
		return true
	}

	return false
}
