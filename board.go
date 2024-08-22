package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) Board() [][]*schemas.Square {
	output := [][]*schemas.Square{}
	row := []*schemas.Square{}

	for i := constants.Ox88["a8"]; i <= constants.Ox88["h1"]; i++ {
		if c.board[i] == nil {
			row = append(row, nil)
		} else {
			row = append(row, &schemas.Square{
				Coords:    utils.Ox88ToSan(i),
				PieceType: c.board[i].PieceType,
				Color:     c.board[i].Color,
			})
		}

		if ((i + 1) & 0x88) != 0 {
			output = append(output, row)
			row = []*schemas.Square{}
			i += 8
		}
	}

	return output
}
