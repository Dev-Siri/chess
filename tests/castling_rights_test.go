package tests

import (
	"testing"

	"github.com/Dev-Siri/chess"
	"github.com/Dev-Siri/chess/constants"
)

func TestCastlingRightsClearWhiteKingSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsClearWhiteKingSide. %v", err)
	}

	if !game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceKing: false,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == true")
	}

	if game.GetCastlingRights(constants.ColorWhite).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('w').KingSide == false")
	}
}

func TestCastlingRightsClearWhiteQueenSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsClearWhiteQueenSide. %v", err)
	}

	if !game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceQueen: false,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == true")
	}

	if game.GetCastlingRights(constants.ColorWhite).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('w').QueenSide == false")
	}
}

func TestCastlingRightsClearBlackKingSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsClearBlackKingSide. %v", err)
	}

	if !game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceKing: false,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == true")
	}

	if game.GetCastlingRights(constants.ColorBlack).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('b').KingSide == false")
	}
}

func TestCastlingRightsClearBlackQueenSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsClearBlackQueenSide. %v", err)
	}

	if !game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceQueen: false,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == true")
	}

	if game.GetCastlingRights(constants.ColorBlack).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('b').QueenSide == false")
	}
}

const positionFen = "r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"

func TestCastlingRightsSetWhiteKingSide(t *testing.T) {
	const positionFen = "r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsSetWhiteKingSide from position(%s). %v", positionFen, err)
	}

	if !game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceKing: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == true")
	}

	if !game.GetCastlingRights(constants.ColorWhite).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('w').KingSide == true")
	}
}

func TestCastlingRightsSetWhiteQueenSide(t *testing.T) {
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsSetWhiteQueenSide from position(%s). %v", positionFen, err)
	}

	if !game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceQueen: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == true")
	}

	if !game.GetCastlingRights(constants.ColorWhite).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('w').QueenSide == true")
	}
}

func TestCastlingRightsSetBlackKingSide(t *testing.T) {
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsSetBlackKingSide from position(%s). %v", positionFen, err)
	}

	if !game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceKing: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == true")
	}

	if !game.GetCastlingRights(constants.ColorBlack).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('b').KingSide == true")
	}
}

func TestCastlingRightsSetBlackQueenSide(t *testing.T) {
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsSetBlackQueenSide from position(%s). %v", positionFen, err)
	}

	if !game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceQueen: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == true")
	}

	if !game.GetCastlingRights(constants.ColorBlack).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('b').QueenSide == true")
	}
}

func TestCastlingRightsFailToSetWhiteKingSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsFailToSetWhiteKingSide. %v", err)
	}

	game.Clear(false)

	if game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceKing: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == false")
	}

	if game.GetCastlingRights(constants.ColorWhite).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('w').KingSide == false")
	}
}

func TestCastlingRightsFailToSetWhiteQueenSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsFailToSetWhiteQueenSide. %v", err)
	}

	game.Clear(false)

	if game.SetCastlingRights(constants.ColorWhite, map[string]bool{
		constants.PieceQueen: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('w', ...) == false")
	}

	if game.GetCastlingRights(constants.ColorWhite).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('w').QueenSide == false")
	}
}

func TestCastlingRightsFailToSetBlackKingSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsFailToSetBlackKingSide. %v", err)
	}

	game.Clear(false)

	if game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceKing: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == false")
	}

	if game.GetCastlingRights(constants.ColorBlack).KingSide {
		t.Fatal("Expected chess.GetCastlingRights('b').KingSide == false")
	}
}

func TestCastlingRightsFailToSetBlackQueenSide(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestCastlingRightsFailToSetBlackQueenSide. %v", err)
	}

	game.Clear(false)

	if game.SetCastlingRights(constants.ColorBlack, map[string]bool{
		constants.PieceQueen: true,
	}) {
		t.Fatal("Expected chess.SetCastlingRights('b', ...) == false")
	}

	if game.GetCastlingRights(constants.ColorBlack).QueenSide {
		t.Fatal("Expected chess.GetCastlingRights('b').QueenSide == false")
	}
}
