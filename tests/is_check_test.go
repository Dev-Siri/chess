package tests

import (
	"testing"

	"github.com/Dev-Siri/chess"
)

func TestIsCheckStartingPosition(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsCheckStartingPosition. %v", err)
		return
	}

	if game.IsCheck() {
		t.Fatalf("Expected game.IsCheck() == false")
	}
}

func TestIsCheckFromPosition(t *testing.T) {
	const positionFen = "rnb1kbnr/pppp1ppp/8/8/4Pp1q/2N5/PPPP2PP/R1BQKBNR w KQkq - 2 4"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsCheckFromPosition from position(%s). %v", positionFen, err)
		return
	}

	if !game.IsCheck() {
		t.Fatalf("Expected game.IsCheck() == true")
	}
}

func TestIsCheckCheckmateIsCheck(t *testing.T) {
	const positionFen = "R3k3/8/4K3/8/8/8/8/8 b - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsCheckCheckmateIsCheck from position(%s). %v", positionFen, err)
		return
	}

	if !game.IsCheck() {
		t.Fatalf("Expected game.IsCheck() == true")
	}
}

func TestIsCheckStalemateIsntCheck(t *testing.T) {
	const positionFen = "4k3/4P3/4K3/8/8/8/8/8 b - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestIsCheckStalemateIsntCheck from position(%s). %v", positionFen, err)
		return
	}

	if game.IsCheck() {
		t.Fatalf("Expected game.IsCheck() == false")
	}
}
