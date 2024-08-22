package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) GetCastlingRights(color string) *schemas.CastlingRights {
	return &schemas.CastlingRights{
		KingSide:  (c.castling[color] & constants.Sides[constants.PieceKing]) != 0,
		QueenSide: (c.castling[color] & constants.Sides[constants.PieceQueen]) != 0,
	}
}

func (c *Chess) SetCastlingRights(color string, rights map[string]bool) bool {
	for _, side := range []string{constants.PieceKing, constants.PieceQueen} {
		if rightAllowed, exists := rights[side]; exists {
			if rightAllowed {
				c.castling[color] |= constants.Sides[side]
			} else {
				c.castling[color] &= ^constants.Sides[side]
			}
		}
	}

	c.updateCastlingRights()
	result := c.GetCastlingRights(color)

	_, kingRightExists := rights[constants.PieceKing]
	_, queenRightExists := rights[constants.PieceQueen]

	return ((!kingRightExists || (rights[constants.PieceKing] == result.KingSide)) && (!queenRightExists || (rights[constants.PieceQueen] == result.QueenSide)))
}
