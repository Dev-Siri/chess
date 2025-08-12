package chess

import (
	"fmt"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

func strip(result *[]string) bool {
	if len(*result) > 0 && (*result)[len(*result)-1] == " " {
		*result = (*result)[:len(*result)-1]
		return true
	}

	return false
}

func wrapComment(width int, move string, result []string, newline string, maxWidth int) int {
	for _, token := range strings.Split(move, "") {
		if token == "" {
			continue
		}

		if (width + len(token)) > maxWidth {
			for strip(&result) {
				width--
			}

			result = append(result, newline)
			width = 0
		}

		result = append(result, token)
		width += len(token)
		result = append(result, " ")
		width++
	}

	if strip(&result) {
		width--
	}

	return width
}

func (c *Chess) Pgn(newline string, maxWidth int) string {
	if newline == "" {
		newline = "\n"
	}

	result := []string{}
	headerExists := false

	for key := range c.header {
		result = append(result, ("[" + key + ` "` + c.header[key] + `"]` + newline))
		headerExists = true
	}

	if headerExists && len(c.history) > 0 {
		result = append(result, newline)
	}

	appendComment := func(moveString string) string {
		comment, commentExists := c.comments[c.Fen()]

		if commentExists {
			var delimiter string

			if len(moveString) > 0 {
				delimiter = " "
			}

			moveString = moveString + delimiter + "{" + comment + "}"
		}

		return moveString
	}

	reversedHistory := []schemas.InternalMove{}

	for len(c.history) > 0 {
		reversedHistory = append(reversedHistory, *c.undoMove())
	}

	moves := []string{}
	var moveString string

	if len(reversedHistory) == 0 {
		moves = append(moves, appendComment(""))
	}

	for len(reversedHistory) > 0 {
		moveString = appendComment(moveString)
		move := &schemas.InternalMove{}

		if len(reversedHistory) != 0 {
			move = &reversedHistory[len(reversedHistory)-1]
			reversedHistory = reversedHistory[:len(reversedHistory)-1]
		}

		if move == nil {
			break
		}

		if len(c.history) < 1 && move.Color == constants.ColorBlack {
			prefix := fmt.Sprintf("%v. ...", c.moveNumber)

			if moveString != "" {
				moveString = moveString + " " + prefix
				moveString = prefix
			} else if move.Color == constants.ColorWhite {
				if len(moveString) > 0 {
					moves = append(moves, moveString)
				}

				moveString = fmt.Sprintf("%v.", c.moveNumber)
			}

			moveString = moveString + " " + c.moveToSan(move, c.moves(true, "", ""))
			c.makeMove(move)
		}

		if len(moveString) > 0 {
			moves = append(moves, appendComment(moveString))
		}

		resultHeader, exists := c.header["Result"]
		if exists {
			moves = append(moves, resultHeader)
		}

		if maxWidth == 0 {
			return strings.Join(result, "") + strings.Join(moves, " ")
		}

		var currentWidth int

		for i := 0; i < len(moves); i++ {
			if (currentWidth + len(moves[i])) > maxWidth {
				if strings.Contains(moves[i], "{") {
					currentWidth = wrapComment(currentWidth, moves[i], result, newline, maxWidth)
					continue
				}
			}

			if ((currentWidth + len(moves[i])) > maxWidth) && (i != 0) {
				if result[len(result)-1] == " " {
					result = result[:len(result)-1]
				}

				result = append(result, newline)
				currentWidth = 0
			} else if i != 0 {
				result = append(result, " ")
				currentWidth++
			}

			result = append(result, moves[i])
			currentWidth += len(moves[i])
		}

	}

	return strings.Join(result, "")
}
