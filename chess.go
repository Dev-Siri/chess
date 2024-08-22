package chess

import (
	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/schemas"
)

type Chess struct {
	board         [128]*schemas.Piece
	turn          string
	header        map[string]string
	kings         map[string]int
	epSquare      int
	halfMoves     int
	moveNumber    int
	history       []schemas.History
	comments      map[string]string
	castling      map[string]int
	positionCount map[string]int
}

func createNewGameInstance() *Chess {
	return &Chess{
		turn:   constants.ColorWhite,
		header: make(map[string]string),
		kings: map[string]int{
			constants.ColorWhite: constants.EmptySquare,
			constants.ColorBlack: constants.EmptySquare,
		},
		epSquare:   constants.EmptySquare,
		halfMoves:  0,
		moveNumber: 0,
		history:    []schemas.History{},
		comments:   make(map[string]string),
		castling: map[string]int{
			constants.ColorWhite: 0,
			constants.ColorBlack: 0,
		},
		positionCount: make(map[string]int),
	}
}

func NewGame() (*Chess, error) {
	chess := createNewGameInstance()

	if err := chess.Load(constants.DefaultPositionFen, true, false); err != nil {
		return nil, err
	}

	return chess, nil
}

func NewGameFromFen(fen string) (*Chess, error) {
	chess := createNewGameInstance()

	if err := chess.Load(fen, false, false); err != nil {
		return nil, err
	}

	return chess, nil
}
