package chess

import (
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) Ascii() string {
	s := "   +------------------------+\n"

	for i := constants.Ox88["a8"]; i <= constants.Ox88["h1"]; i++ {
		if utils.File(i) == 0 {
			s += " " + string("87654321"[utils.Rank(i)]) + " |"
		}

		if c.board[i] != nil {
			piece := c.board[i].PieceType
			color := c.board[i].Color

			var symbol string

			if color == constants.ColorWhite {
				symbol = utils.LetterToChessUnicode(strings.ToUpper(piece))
			} else {
				symbol = utils.LetterToChessUnicode(strings.ToLower(piece))
			}

			s += " " + symbol + " "
		} else {
			s += " . "
		}

		if ((i + 1) & 0x88) != 0 {
			s += "|\n"
			i += 8
		}
	}

	s += "   +------------------------+\n"
	s += "     a  b  c  d  e  f  g  h"

	return s
}
