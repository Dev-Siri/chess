package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) SquareColor(square string) string {
	if _, exists := constants.Ox88[square]; exists {
		sq := constants.Ox88[square]

		if ((utils.Rank(sq) + utils.File(sq)) % 2) == 0 {
			return "light"
		} else {
			return "dark"
		}
	}

	return ""
}
