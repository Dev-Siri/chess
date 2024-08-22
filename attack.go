package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) attacked(color string, square int) (bool, []string) {
	attackers := []string{}
	isAttacked := false

	for i := constants.Ox88["a8"]; i <= constants.Ox88["h1"]; i++ {
		// Did we run off the end of the board
		if (i & 0x88) != 0 {
			i += 7
			continue
		}

		// If empty square or wrong color
		if c.board[i] == nil || c.board[i].Color != color {
			continue
		}

		piece := c.board[i]
		difference := i - square

		// Skip - to/from square are the same
		if difference == 0 {
			continue
		}

		index := difference + 119

		if (constants.Attacks[index] & constants.PieceMasks[piece.PieceType]) != 0 {
			if piece.PieceType == constants.PiecePawn {
				if (difference > 0 && piece.Color == constants.ColorWhite) ||
					(difference <= 0 && piece.Color == constants.ColorBlack) {
					isAttacked = true
					attackers = append(attackers, utils.Ox88ToSan(i))
				}

				continue
			}

			// If the piece is a Knight or a King
			if piece.PieceType == constants.PieceKnight || piece.PieceType == constants.PieceKing {
				isAttacked = true
				attackers = append(attackers, utils.Ox88ToSan(i))
				continue
			}

			offset := constants.Rays[index]
			j := i + offset

			blocked := false

			for j != square {
				if c.board[j] != nil {
					blocked = true
					break
				}

				j += offset
			}

			if !blocked {
				isAttacked = true
				attackers = append(attackers, utils.Ox88ToSan(i))
				continue
			}
		}
	}

	return isAttacked, attackers
}

func (c *Chess) isKingAttacked(color string) bool {
	square := c.kings[color]

	if square == -1 {
		return false
	}

	isAttacked, _ := c.attacked(utils.SwapColors(color), square)

	return isAttacked
}

func (c *Chess) Attackers(square string, attackedBy string) []string {
	if attackedBy == "" {
		_, attackers := c.attacked(c.turn, constants.Ox88[square])
		return attackers
	}

	_, attackers := c.attacked(attackedBy, constants.Ox88[square])
	return attackers
}

func (c *Chess) IsAttacked(square string, attackedBy string) bool {
	isAttacked, _ := c.attacked(attackedBy, constants.Ox88[square])
	return isAttacked
}
