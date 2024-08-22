package tests

import (
	"strings"
	"testing"

	"github.com/Dev-Siri/chess"
)

func TestAscii(t *testing.T) {
	output := []string{
		"   +------------------------+",
		" 8 | ♜  .  .  .  .  ♜  ♚  . |",
		" 7 | .  .  .  .  ♞  ♛  ♟  ♟ |",
		" 6 | .  ♟  .  ♟  .  .  .  . |",
		" 5 | .  .  ♟  ♙  ♟  ♟  .  . |",
		" 4 | ♝  ♙  ♙  .  ♙  .  .  . |",
		" 3 | ♖  .  ♗  .  ♘  ♕  .  . |",
		" 2 | ♙  .  .  .  .  ♙  ♙  ♙ |",
		" 1 | .  ♖  .  .  .  .  ♔  . |",
		"   +------------------------+",
		"     a  b  c  d  e  f  g  h",
	}

	const startingFen = "r4rk1/4nqpp/1p1p4/2pPpp2/bPP1P3/R1B1NQ2/P4PPP/1R4K1 w - - 0 28"

	game, err := chess.NewGameFromFen(startingFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAscii.\nchess.NewGameFromFen(%s) = %v", startingFen, err)
		return
	}

	asciiResult := game.Ascii()
	expectedResult := strings.Join(output, "\n")

	if asciiResult != expectedResult {
		t.Fatalf("chess.Ascii() = %s, expected %s", asciiResult, expectedResult)
	}
}
