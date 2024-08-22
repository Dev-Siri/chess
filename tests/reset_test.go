package tests

import (
	"testing"

	"github.com/Dev-Siri/chess"
	"github.com/Dev-Siri/chess/constants"
)

func TestReset(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestReset. %v", err)
	}

	game.Clear(false)
	game.Reset()

	if game.Fen() != constants.DefaultPositionFen {
		t.Fatalf("Expected game.Fen() == %s", constants.DefaultPositionFen)
	}
}
