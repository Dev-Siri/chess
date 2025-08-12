package utils

import (
	"regexp"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

// Extracts the zero-based rank of an 0x88 square
func Rank(square int) int {
	return square >> 4
}

// Extracts the zero-based file of an 0x88 square
func File(square int) int {
	return square & 0xf
}

// Converts a 0x88 value (int) to Standard Algebraic Notation
func Ox88ToSan(square int) string {
	file := File(square)
	rank := Rank(square)
	san := "abcdefgh"[file:file+1] + "87654321"[rank:rank+1]

	return san
}

// Uniquely identifies ambiguous moves
func GetDisambiguator(move *schemas.InternalMove, moves []schemas.InternalMove) string {
	from, to, piece := move.From, move.To, move.Piece

	ambiguities, sameRank, sameFile := 0, 0, 0

	for i := 0; i < len(moves); i++ {
		ambigFrom, ambigTo, ambigPiece := moves[i].From, moves[i].To, moves[i].Piece

		// If a move of the same piece type ends on the same `to` square, we'll need to
		// add a disambiguator to the Algebraic Notation
		if piece == ambigPiece && from != ambigFrom && to == ambigTo {
			ambiguities++

			if Rank(from) == Rank(ambigFrom) {
				sameRank++
			}

			if File(from) == File(ambigFrom) {
				sameFile++
			}
		}
	}

	if ambiguities > 0 {
		if sameRank > 0 && sameFile > 0 {
			// If there exists a similar moving piece on the same rank and file as
			// the move in question, use the square as the disambiguator
			return Ox88ToSan(from)
		} else if sameFile > 0 {
			// If the moving piece rests on the same file, use the rank symbol as the disambiguator
			return string(Ox88ToSan(from)[1])
		} else {
			// Else use the file symbol
			return string(Ox88ToSan(from)[0])
		}
	}

	return ""
}

func AddMove(moves *[]schemas.InternalMove, color string, from int, to int, piece string, captured string, flags int) {
	if flags == 0 {
		flags = constants.Bits[constants.BitNormal]
	}

	rank := Rank(to)

	if piece == constants.PiecePawn && (rank == constants.Rank1 || rank == constants.Rank8) {
		for i := 0; i < len(constants.Promotions); i++ {
			promotion := constants.Promotions[i]
			*moves = append(*moves, schemas.InternalMove{
				Color:     color,
				From:      from,
				To:        to,
				Piece:     piece,
				Captured:  captured,
				Promotion: promotion,
				Flags:     flags | constants.Bits[promotion],
			})
		}
	} else {
		*moves = append(*moves, schemas.InternalMove{
			Color:    color,
			From:     from,
			To:       to,
			Piece:    piece,
			Captured: captured,
			Flags:    flags,
		})
	}
}

func PopMove(history []schemas.History) ([]schemas.History, schemas.History) {
	if len(history) == 0 {
		return history, schemas.History{}
	}
	lastElem := history[len(history)-1]
	history = history[:len(history)-1]

	return history, lastElem
}

func InferPieceType(san string) string {
	pieceType := san[0]
	pieceTypeStr := string(pieceType)

	if pieceType >= 'a' && pieceType <= 'h' {
		re := regexp.MustCompile(`[a-h]\d.*[a-h]\d`)

		if re.MatchString(san) {
			return ""
		}

		return constants.PiecePawn
	}

	pieceTypeLower := strings.ToLower(string(pieceTypeStr))

	if pieceTypeLower == "o" {
		return constants.PieceKing
	}

	return pieceTypeLower
}

// Parses all of the decorators out of a SAN string
func StrippedSan(move string) string {
	promotionRe := regexp.MustCompile("=")
	moveQualityRe := regexp.MustCompile("[+#]?[?!]*$")

	return moveQualityRe.ReplaceAllString(promotionRe.ReplaceAllString(move, ""), "")
}

// Remove the last two fields in FEN string as they're not needed when checking for repetition
func TrimFen(fen string) string {
	return strings.Join(strings.Split(fen, " ")[0:4], " ")
}
