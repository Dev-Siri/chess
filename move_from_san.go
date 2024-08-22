package chess

import (
	"regexp"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
	"github.com/Dev-Siri/chess/utils"
)

func (c *Chess) moveFromSan(move string, strict bool) *schemas.InternalMove {
	cleanMove := utils.StrippedSan(move)

	pieceType := utils.InferPieceType(cleanMove)
	moves := c.moves(true, pieceType, "")

	for i := 0; i < len(moves); i++ {
		if cleanMove == utils.StrippedSan(c.moveToSan(&moves[i], moves)) {
			return &moves[i]
		}
	}

	if strict {
		return nil
	}

	piece, from, to, promotion := "", "", "", ""
	matchesRe := regexp.MustCompile("([pnbrqkPNBRQK])?([a-h][1-8])x?-?([a-h][1-8])([qrbnQRBN])?")
	matches := matchesRe.FindAllString(cleanMove, -1)
	overlyDisambiguated := false

	if len(matches) > 0 {
		piece = matches[1]
		from = matches[2]
		to = matches[3]
		promotion = matches[4]

		if len(from) == 1 {
			overlyDisambiguated = true
		}
	} else {
		matchesRe := regexp.MustCompile("([pnbrqkPNBRQK])?([a-h]?[1-8]?)x?-?([a-h][1-8])([qrbnQRBN])?")
		matches = matchesRe.FindAllString(cleanMove, -1)

		if len(matches) > 0 {
			piece = matches[1]
			from = matches[2]
			to = matches[3]
			promotion = matches[4]

			if len(from) == 1 {
				overlyDisambiguated = true
			}
		}
	}

	pieceType = utils.InferPieceType(cleanMove)

	if piece != "" {
		moves = c.moves(true, piece, "")
	} else {
		moves = c.moves(true, pieceType, "")
	}

	if to == "" {
		return nil
	}

	for i := 0; i < len(moves); i++ {
		if from == "" {
			if cleanMove == strings.ReplaceAll(utils.StrippedSan(c.moveToSan(&moves[i], moves)), "x", "") {
				return &moves[i]
			}
		} else if (piece == "" || strings.ToLower(piece) == moves[i].Piece) &&
			constants.Ox88[from] == moves[i].From && constants.Ox88[to] == moves[i].To &&
			(promotion == "" || strings.ToLower(promotion) == moves[i].Promotion) {
			return &moves[i]
		} else if overlyDisambiguated {
			square := utils.Ox88ToSan(moves[i].From)

			if (piece == "" && strings.ToLower(piece) == moves[i].Piece) &&
				constants.Ox88[to] == moves[i].To && (from == string(square[0]) || from == string(square[1])) &&
				(promotion == "" || strings.ToLower(promotion) == moves[i].Promotion) {
				return &moves[i]
			}
		}
	}

	return nil
}
