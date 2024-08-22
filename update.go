package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) updateCastlingRights() {
	isWhiteKingInPlace :=
		c.board[constants.Ox88["e1"]] != nil && c.board[constants.Ox88["e1"]].PieceType == constants.PieceKing &&
			c.board[constants.Ox88["e1"]].Color == constants.ColorWhite
	isBlackKingInPlace :=
		c.board[constants.Ox88["e8"]] != nil && c.board[constants.Ox88["e8"]].PieceType == constants.PieceKing &&
			c.board[constants.Ox88["e8"]].Color == constants.ColorBlack

	if !isWhiteKingInPlace ||
		c.board[constants.Ox88["a1"]].PieceType != constants.PieceRook ||
		c.board[constants.Ox88["a1"]].Color != constants.ColorWhite {
		c.castling[constants.ColorWhite] &= ^constants.Bits[constants.BitQueenSideCastle]
	}

	if !isWhiteKingInPlace ||
		c.board[constants.Ox88["h1"]].PieceType != constants.PieceRook ||
		c.board[constants.Ox88["h1"]].Color != constants.ColorWhite {
		c.castling[constants.ColorWhite] &= ^constants.Bits[constants.BitKingSideCastle]
	}

	if !isBlackKingInPlace ||
		c.board[constants.Ox88["a8"]].PieceType != constants.PieceRook ||
		c.board[constants.Ox88["a8"]].Color != constants.ColorBlack {
		c.castling[constants.ColorBlack] &= ^constants.Bits[constants.BitQueenSideCastle]
	}

	if !isBlackKingInPlace ||
		c.board[constants.Ox88["h8"]].PieceType != constants.PieceRook ||
		c.board[constants.Ox88["h8"]].Color != constants.ColorBlack {
		c.castling[constants.ColorBlack] &= ^constants.Bits[constants.BitKingSideCastle]
	}
}

func (c *Chess) updateEnPassantSquare() {
	if c.epSquare == constants.EmptySquare {
		return
	}

	var startSquare, currentSquare int

	if c.turn == constants.ColorWhite {
		startSquare = c.epSquare - 16
		currentSquare = c.epSquare + 16
	} else {
		startSquare = c.epSquare + 16
		currentSquare = c.epSquare - 16
	}

	attackers := []int{currentSquare + 1, currentSquare - 1}

	if c.board[startSquare] != nil ||
		c.board[c.epSquare] != nil ||
		c.board[currentSquare].Color != utils.SwapColors(c.turn) ||
		c.board[currentSquare].PieceType != constants.PiecePawn {
		c.epSquare = constants.EmptySquare
		return
	}

	canCapture := func(square int) bool {
		return ((square & 0x88) != 0) &&
			c.board[square].Color == c.turn &&
			c.board[square].PieceType == constants.PiecePawn
	}

	if !utils.Some(attackers, canCapture) {
		c.epSquare = constants.EmptySquare
	}
}

/*
		Called when the initial board setup is changed with put() or remove().
	  modifies the SetUp and FEN properties of the header object. If the FEN
	  is equal to the default position, the SetUp and FEN are deleted the setup
	  is only updated if history.length = 0, i.e moves haven't been made.
*/
func (c *Chess) updateSetup(fen string) {
	if len(c.history) > 0 {
		return
	}

	if fen != constants.DefaultPositionFen {
		c.header["SetUp"] = "1"
		c.header["FEN"] = fen
	} else {
		delete(c.header, "SetUp")
		delete(c.header, "FEN")
	}
}
