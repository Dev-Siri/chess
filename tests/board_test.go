package tests

import (
	"reflect"
	"testing"

	"github.com/Dev-Siri/chess"
	"github.com/Dev-Siri/chess/schemas"
)

type BoardTest struct {
	Fen   string
	Board [][]*schemas.Square
}

func areBoardsEqual(board1, board2 [][]*schemas.Square) bool {
	if len(board1) != len(board2) {
		return false
	}

	for i := range board1 {
		if len(board1[i]) != len(board2[i]) {
			return false
		}

		for j := range board1[i] {
			if !reflect.DeepEqual(board1[i][j], board2[i][j]) {
				return false
			}
		}
	}

	return true
}

var tests = []BoardTest{
	{
		Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		Board: [][]*schemas.Square{
			{
				{Coords: "a8", PieceType: "r", Color: "b"},
				{Coords: "b8", PieceType: "n", Color: "b"},
				{Coords: "c8", PieceType: "b", Color: "b"},
				{Coords: "d8", PieceType: "q", Color: "b"},
				{Coords: "e8", PieceType: "k", Color: "b"},
				{Coords: "f8", PieceType: "b", Color: "b"},
				{Coords: "g8", PieceType: "n", Color: "b"},
				{Coords: "h8", PieceType: "r", Color: "b"},
			},
			{
				{Coords: "a7", PieceType: "p", Color: "b"},
				{Coords: "b7", PieceType: "p", Color: "b"},
				{Coords: "c7", PieceType: "p", Color: "b"},
				{Coords: "d7", PieceType: "p", Color: "b"},
				{Coords: "e7", PieceType: "p", Color: "b"},
				{Coords: "f7", PieceType: "p", Color: "b"},
				{Coords: "g7", PieceType: "p", Color: "b"},
				{Coords: "h7", PieceType: "p", Color: "b"},
			},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{
				{Coords: "a2", PieceType: "p", Color: "w"},
				{Coords: "b2", PieceType: "p", Color: "w"},
				{Coords: "c2", PieceType: "p", Color: "w"},
				{Coords: "d2", PieceType: "p", Color: "w"},
				{Coords: "e2", PieceType: "p", Color: "w"},
				{Coords: "f2", PieceType: "p", Color: "w"},
				{Coords: "g2", PieceType: "p", Color: "w"},
				{Coords: "h2", PieceType: "p", Color: "w"},
			},
			{
				{Coords: "a1", PieceType: "r", Color: "w"},
				{Coords: "b1", PieceType: "n", Color: "w"},
				{Coords: "c1", PieceType: "b", Color: "w"},
				{Coords: "d1", PieceType: "q", Color: "w"},
				{Coords: "e1", PieceType: "k", Color: "w"},
				{Coords: "f1", PieceType: "b", Color: "w"},
				{Coords: "g1", PieceType: "n", Color: "w"},
				{Coords: "h1", PieceType: "r", Color: "w"},
			},
		},
	},
	// Checkmate
	{
		Fen: "r3k2r/ppp2p1p/2n1p1p1/8/2B2P1q/2NPb1n1/PP4PP/R2Q3K w kq - 0 8",
		Board: [][]*schemas.Square{
			{
				{Coords: "a8", PieceType: "r", Color: "b"},
				nil,
				nil,
				nil,
				{Coords: "e8", PieceType: "k", Color: "b"},
				nil,
				nil,
				{Coords: "h8", PieceType: "r", Color: "b"},
			},
			{
				{Coords: "a7", PieceType: "p", Color: "b"},
				{Coords: "b7", PieceType: "p", Color: "b"},
				{Coords: "c7", PieceType: "p", Color: "b"},
				nil,
				nil,
				{Coords: "f7", PieceType: "p", Color: "b"},
				nil,
				{Coords: "h7", PieceType: "p", Color: "b"},
			},
			{
				nil,
				nil,
				{Coords: "c6", PieceType: "n", Color: "b"},
				nil,
				{Coords: "e6", PieceType: "p", Color: "b"},
				nil,
				{Coords: "g6", PieceType: "p", Color: "b"},
				nil,
			},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{
				nil,
				nil,
				{Coords: "c4", PieceType: "b", Color: "w"},
				nil,
				nil,
				{Coords: "f4", PieceType: "p", Color: "w"},
				nil,
				{Coords: "h4", PieceType: "q", Color: "b"},
			},
			{
				nil,
				nil,
				{Coords: "c3", PieceType: "n", Color: "w"},
				{Coords: "d3", PieceType: "p", Color: "w"},
				{Coords: "e3", PieceType: "b", Color: "b"},
				nil,
				{Coords: "g3", PieceType: "n", Color: "b"},
				nil,
			},
			{
				{Coords: "a2", PieceType: "p", Color: "w"},
				{Coords: "b2", PieceType: "p", Color: "w"},
				nil,
				nil,
				nil,
				nil,
				{Coords: "g2", PieceType: "p", Color: "w"},
				{Coords: "h2", PieceType: "p", Color: "w"},
			},
			{
				{Coords: "a1", PieceType: "r", Color: "w"},
				nil,
				nil,
				{Coords: "d1", PieceType: "q", Color: "w"},
				nil,
				nil,
				nil,
				{Coords: "h1", PieceType: "k", Color: "w"},
			},
		},
	},
}

func TestBoard(t *testing.T) {
	for _, test := range tests {
		game, err := chess.NewGameFromFen(test.Fen)

		if err != nil {
			t.Fatalf("Failed to initialize game for test TestBoard for board position %s. %v", test.Fen, err)
		}

		if !areBoardsEqual(game.Board(), test.Board) {
			t.Fatal("game.Board() != test.Board, expected equality")
		}
	}
}
