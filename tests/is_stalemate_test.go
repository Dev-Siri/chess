package tests

import (
	"testing"

	"github.com/Dev-Siri/chess"
)

func TestIsStalemateFirstPosition(t *testing.T) {
	const positionFen = "1R6/8/8/8/8/8/7R/k6K b - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsStalemateFirstPosition from position(%s). %v", positionFen, err)
		return
	}

	if !game.IsStalemate() {
		t.Fatalf("Expected game.IsStalemate() == true")
	}
}

func TestIsStalemateSecondPosition(t *testing.T) {
	const positionFen = "8/8/5k2/p4p1p/P4K1P/1r6/8/8 w - - 0 2"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsStalemateSecondPosition from position(%s). %v", positionFen, err)
		return
	}

	if !game.IsStalemate() {
		t.Fatalf("Expected game.IsStalemate() == true")
	}
}

func TestIsStalemateStartingPosition(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsStalemateSecondPosition. %v", err)
		return
	}

	if game.IsStalemate() {
		t.Fatalf("Expected game.IsStalemate() == false")
	}
}

func TestIsStalemateCheckmateIsntStalemate(t *testing.T) {
	const positionFen = "R3k3/8/4K3/8/8/8/8/8 b - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsStalemateCheckmateIsntStalemate from position(%s). %v", positionFen, err)
		return
	}

	if game.IsStalemate() {
		t.Fatalf("Expected game.IsStalemate() == false")
	}
}
