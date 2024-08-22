package chess

import (
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) moveToSan(move *schemas.InternalMove, moves []schemas.InternalMove) string {
	var output string

	if (move.Flags & constants.Bits[constants.BitKingSideCastle]) != 0 {
		output = "O-O"
	} else if (move.Flags & constants.Bits[constants.BitQueenSideCastle]) != 0 {
		output = "O-O-O"
	} else {
		if move.Piece != constants.PiecePawn {
			disambiguator := utils.GetDisambiguator(move, moves)
			output += strings.ToUpper(move.Piece) + disambiguator
		}

		if (move.Flags & (constants.Bits[constants.BitCapture] | constants.Bits[constants.BitEpCapture])) != 0 {
			if move.Piece == constants.PiecePawn {
				output += string(utils.Ox88ToSan(move.From)[0])
			}

			output += "x"
		}

		output += utils.Ox88ToSan(move.To)

		if move.Promotion != "" {
			output += "=" + strings.ToUpper(move.Promotion)
		}
	}

	c.makeMove(move)
	if c.IsCheck() {
		if c.IsCheckmate() {
			output += "#"
		} else {
			output += "+"
		}
	}

	c.undoMove()

	return output
}
