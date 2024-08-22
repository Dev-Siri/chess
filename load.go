package chess

import (
	"strconv"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) Load(fen string, skipValidation bool, preserveHeaders bool) error {
	tokens := strings.Fields(fen)
	editableFen := fen

	// Append commonly omitted fen tokens
	if len(tokens) >= 2 && len(tokens) < 6 {
		adjustments := []string{"-", "-", "0", "1"}
		missingTokensCount := 6 - len(tokens)

		editableFen = strings.Join(append(tokens, adjustments[len(adjustments)-missingTokensCount:]...), " ")
	}

	tokens = strings.Fields(editableFen)

	if !skipValidation {
		if err := ValidateFen(editableFen); err != nil {
			return err
		}
	}

	position := tokens[0]
	square := 0

	c.Clear(preserveHeaders)

	for i := 0; i < len(position); i++ {
		piece := position[i]
		pieceStr := string(piece)

		if pieceStr == "/" {
			square += 8
		} else if utils.IsDigit(pieceStr) {
			pieceInt, err := strconv.Atoi(pieceStr)

			if err != nil {
				return err
			}

			square += pieceInt
		} else {
			color := ""

			if piece < 'a' {
				color = constants.ColorWhite
			} else {
				color = constants.ColorBlack
			}

			c.put(strings.ToLower(pieceStr), color, utils.Ox88ToSan(square))

			square++
		}
	}

	c.turn = tokens[1]

	if strings.Contains(tokens[2], "K") {
		c.castling[constants.ColorWhite] |= constants.Bits[constants.BitKingSideCastle]
	}

	if strings.Contains(tokens[2], "Q") {
		c.castling[constants.ColorWhite] |= constants.Bits[constants.BitQueenSideCastle]
	}

	if strings.Contains(tokens[2], "k") {
		c.castling[constants.ColorBlack] |= constants.Bits[constants.BitKingSideCastle]
	}

	if strings.Contains(tokens[2], "q") {
		c.castling[constants.ColorBlack] |= constants.Bits[constants.BitQueenSideCastle]
	}

	eps := string(tokens[3])

	if eps == "-" {
		c.epSquare = constants.EmptySquare
	} else {
		c.epSquare = constants.Ox88[eps]
	}

	halfMovesInt, err := strconv.Atoi(string(tokens[4]))

	if err != nil {
		return err
	}

	moveNumberInt, err := strconv.Atoi(string(tokens[5]))

	if err != nil {
		return err
	}

	c.halfMoves = halfMovesInt
	c.moveNumber = moveNumberInt

	c.updateSetup(fen)
	c.incPositionCount(fen)

	return nil
}
