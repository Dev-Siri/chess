package chess

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func ValidateFen(fen string) error {
	tokens := strings.Fields(fen)

	// 1st Criterion: 6 space-seperated fields?
	if len(tokens) != 6 {
		return fmt.Errorf("invalid fen: must contain six space-delimated fields")
	}

	// 2nd Criterion: Move number field is a integer value > 0?
	moveNumber, err := strconv.Atoi(tokens[5])

	if err != nil {
		return fmt.Errorf("invalid fen: move number is not parsable as int")
	}

	if moveNumber <= 0 {
		return fmt.Errorf("invalid fen: move number must be a positive int")
	}

	// 3rd Criterion: Half move counter is an integer >= 0?
	halfMoves, err := strconv.Atoi(tokens[4])

	if err != nil {
		return fmt.Errorf("invalid fen: half move counter number is not parsable as int")
	}

	if halfMoves < 0 {
		return fmt.Errorf("invalid fen: half move counter number must be a non-negative number")
	}

	// 4th Criterion: 4th field is a valid e.p.-string?
	enPassantSquareRe := regexp.MustCompile("^(-|[abcdefgh][36])$")
	enPassantSquare := enPassantSquareRe.MatchString(tokens[3])

	if !enPassantSquare {
		return fmt.Errorf("invalid fen: en-passant square is invalid")
	}

	// 5th Criterion: 3th field is a valid castle-string?
	validCastleStringRe := regexp.MustCompile("[^kKqQ-]")
	validCastleString := validCastleStringRe.MatchString(tokens[2])

	if validCastleString {
		return fmt.Errorf("invalid fen: castling availability is invalid")
	}

	// 6th Criterion: 2nd field is "w" (white) or "b" (black)?
	sideToSideMoveRe := regexp.MustCompile("^(w|b)$")
	sideToSideMove := sideToSideMoveRe.MatchString(tokens[1])

	if !sideToSideMove {
		return fmt.Errorf("invalid fen: side-to-side move is invalid")
	}

	// 7th Criterion: 1st field contains 8 rows?
	rows := strings.Split(tokens[0], "/")

	if len(rows) != 8 {
		return fmt.Errorf("invalid fen: piece data does not contain 8 '/'-delimated rows")
	}

	// 8th Criterion: Every row is valid?
	for i := 0; i < len(rows); i++ {
		sumFields := 0
		previousWasNumber := false

		for k := 0; k < len(rows[i]); k++ {
			// go converts to ascii (int) on indexing strings (picking chars)
			// so string() converts that back to just a stringified digit (if actually a digit)
			pieceData := string(rows[i][k])

			if utils.IsDigit(pieceData) {
				if previousWasNumber {
					return fmt.Errorf("invalid fen: piece data is invalid (consecutive number)")
				}

				pieceDataInt, err := strconv.Atoi(pieceData)

				if err != nil {
					return fmt.Errorf("invalid fen: piece data not parsable as number")
				}

				sumFields += pieceDataInt
				previousWasNumber = true
			} else {
				pieceRe := regexp.MustCompile("^[prnbqkPRNBQK]$")
				isValidPiece := pieceRe.MatchString(pieceData)

				if !isValidPiece {
					return fmt.Errorf("invalid fen: piece data is invalid (invalid piece)")
				}

				sumFields++
				previousWasNumber = false
			}
		}

		if sumFields != 8 {
			return fmt.Errorf("invalid fen: piece data is invalid (too many squares in rank)")
		}
	}

	// Actual square for en-passanting
	enPassantBoardSquare := ""

	if len(tokens[3]) > 1 {
		enPassantBoardSquare = string(tokens[3][1])
	}

	// 9th Criterion: Is en-passant square legal?
	if (enPassantBoardSquare == "3" && tokens[1] == constants.ColorWhite) ||
		(enPassantBoardSquare == "6" && tokens[1] == constants.ColorBlack) {
		return fmt.Errorf("invalid fen: illegal en-passant square")
	}

	// 10th Criterion: Does chess position contain exact two kings?
	kings := []schemas.KingPosition{
		{Color: constants.ColorWhite, Regex: regexp.MustCompile("K")},
		{Color: constants.ColorBlack, Regex: regexp.MustCompile("k")},
	}

	for _, king := range kings {
		if !king.Regex.MatchString(tokens[0]) {
			return fmt.Errorf("invalid fen: missing %s king", king.Color)
		}

		if len(king.Regex.FindAllString(tokens[0], -1)) > 1 {
			return fmt.Errorf("invalid fen: too many %s kings", king.Color)
		}
	}

	// 11th Criterion: Are any pawns on the first or eighth rows?
	pawnsOnEdgeOfBoard := false

	for _, piece := range strings.Split(rows[0], "") {
		if strings.ToUpper(piece) == "P" {
			pawnsOnEdgeOfBoard = true
		}
	}

	for _, piece := range strings.Split(rows[7], "") {
		if strings.ToUpper(piece) == "P" {
			pawnsOnEdgeOfBoard = true
		}
	}

	if pawnsOnEdgeOfBoard {
		return fmt.Errorf("invalid fen: some pawns are on the edge rows")
	}

	return nil
}
