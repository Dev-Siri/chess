package tests

import (
	"testing"

	"github.com/Dev-Siri/chess"
	"github.com/Dev-Siri/chess/constants"
)

func getAttackerCount(c *chess.Chess, color string) []int {
	attackerCounts := make([]int, 64)
	for i := 0; i < 64; i++ {
		attackerCounts[i] = len(c.Attackers(constants.Squares[i], color))
	}
	return attackerCounts
}

func TestAttackersCountInDefaultPosition(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersCountInDefaultPosition. %v", err)
	}

	expectedAttackersForWhite := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		2, 2, 3, 2, 2, 3, 2, 2,
		1, 1, 1, 4, 4, 1, 1, 1,
		0, 1, 1, 1, 1, 1, 1, 0,
	}

	resultingAttackersForWhite := getAttackerCount(game, constants.ColorWhite)

	for i, attackerCount := range resultingAttackersForWhite {
		if attackerCount != expectedAttackersForWhite[i] {
			t.Fatalf("Attackers for white in default position != expectedAttackersForWhite.")
		}
	}

	expectedAttackersForBlack := []int{
		0, 1, 1, 1, 1, 1, 1, 0,
		1, 1, 1, 4, 4, 1, 1, 1,
		2, 2, 3, 2, 2, 3, 2, 2,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	resultingAttackersForBlack := getAttackerCount(game, constants.ColorBlack)

	for i, attackerCount := range resultingAttackersForBlack {
		if attackerCount != expectedAttackersForBlack[i] {
			t.Fatalf("Attackers for black in default position != expectedAttackersForBlack.")
		}
	}
}

func TestAttackersCountInMiddleGame(t *testing.T) {
	// Gujrathi–Firouzja, Round 6
	const positionFen = "r3kb1r/1b3ppp/pqnppn2/1p6/4PBP1/PNN5/1PPQBP1P/2KR3R b kq - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersCountInMiddleGame from position %s. %v", positionFen, err)
	}

	expectedAttackersForWhite := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 2, 0, 0, 0, 1,
		1, 2, 1, 3, 1, 2, 1, 1,
		1, 1, 1, 2, 1, 1, 1, 0,
		1, 1, 2, 3, 3, 1, 3, 0,
		1, 1, 2, 4, 2, 0, 0, 2,
		1, 2, 3, 5, 3, 3, 2, 1,
	}

	resultingAttackersForWhite := getAttackerCount(game, constants.ColorWhite)

	for i, attackerCount := range resultingAttackersForWhite {
		if attackerCount != expectedAttackersForWhite[i] {
			t.Fatalf("Attackers for white in position(%s) != expectedAttackersForWhite.", positionFen)
		}
	}

	expectedAttackersForBlack := []int{
		1, 2, 2, 4, 2, 2, 2, 0,
		3, 1, 1, 2, 3, 1, 1, 2,
		3, 0, 2, 1, 1, 1, 2, 1,
		2, 2, 2, 2, 2, 1, 0, 1,
		1, 1, 1, 2, 1, 0, 1, 0,
		0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	resultingAttackersForBlack := getAttackerCount(game, constants.ColorBlack)

	for i, attackerCount := range resultingAttackersForBlack {
		if attackerCount != expectedAttackersForBlack[i] {
			t.Fatalf("Attackers for black in position(%s) != expectedAttackersForBlack.", positionFen)
		}
	}
}

func TestAttackersWhenAllButOneSquareIsCovered(t *testing.T) {
	const positionFen = "Q4K1k/1Q5p/2Q5/3Q4/4Q3/5Q2/6Q1/7Q w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersWhenAllButOneSquareIsCovered from position %s. %v", positionFen, err)
	}

	expectedAttackersForWhite := []int{
		1, 2, 3, 2, 4, 2, 3, 0,
		2, 2, 2, 3, 3, 4, 3, 3,
		3, 2, 2, 2, 3, 2, 3, 2,
		2, 3, 2, 2, 2, 3, 2, 3,
		3, 2, 3, 2, 2, 2, 3, 2,
		2, 3, 2, 3, 2, 2, 2, 3,
		3, 2, 3, 2, 3, 2, 2, 2,
		2, 3, 2, 3, 2, 3, 2, 1,
	}

	resultingAttackersForWhite := getAttackerCount(game, constants.ColorWhite)

	for i, attackerCount := range resultingAttackersForWhite {
		if attackerCount != expectedAttackersForWhite[i] {
			t.Fatalf("Attackers for white in position(%s) != expectedAttackersForWhite.", positionFen)
		}
	}

	expectedAttackersForBlack := []int{
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	resultingAttackersForBlack := getAttackerCount(game, constants.ColorBlack)

	for i, attackerCount := range resultingAttackersForBlack {
		if attackerCount != expectedAttackersForBlack[i] {
			t.Fatalf("Attackers for black in position(%s) != expectedAttackersForBlack.", positionFen)
		}
	}
}

func TestAttackersReturnValueDependingOnSideToMove(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersReturnValueDependingOnSideToMove. %v", err)
	}

	if !IncludesSameMember(game.Attackers("c3", ""), []string{"b1", "b2", "d2"}) {
		t.Fatalf("Attackers for c3 != {b1, b2, d2}.")
	}

	if len(game.Attackers("c6", "")) != 0 {
		t.Fatalf("Attackers for c6 != {}.")
	}

	if _, err = game.Move("e4", false); err != nil {
		t.Fatalf("Failed to execute move `e4` in game for test TestAttackersReturnValueDependingOnSideToMove. %v", err)
	}

	if len(game.Attackers("c3", "")) != 0 {
		t.Fatalf("Attackers for c3 != {} after 1. e4")
	}

	if !IncludesSameMember(game.Attackers("c6", ""), []string{"b7", "b8", "d7"}) {
		t.Fatalf("Attackers for c6 != {b7, b8, d7} after 1. e4")
	}

	if _, err = game.Move("e5", false); err != nil {
		t.Fatalf("Failed to execute move `e5` in game for test TestAttackersReturnValueDependingOnSideToMove. %v", err)
	}

	if !IncludesSameMember(game.Attackers("c3", ""), []string{"b1", "b2", "d2"}) {
		t.Fatalf("Attackers for c6 != {b7, b8, d7} after 1. e4 e5")
	}

	if len(game.Attackers("c6", "")) != 0 {
		t.Fatalf("Attackers for c6 != {} after 1. e4 e5")
	}
}

func TestAttackersEveryPieceAttackingEmptySquare(t *testing.T) {
	const positionFen = "2b5/4kp2/2r5/3q2n1/8/8/4P3/4K3 w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersEveryPieceAttackingEmptySquare from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("e6", constants.ColorBlack), []string{"c6", "c8", "d5", "e7", "f7", "g5"}) {
		t.Fatalf("Attackers for e6 != {c6, c8, d5, e7, f7, g5} in position(%s)", positionFen)
	}
}

func TestAttackersEveryPieceAttackingAnotherPiece(t *testing.T) {
	const positionFen = "4k3/8/8/8/5Q2/5p1R/4PK2/4N2B w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersEveryPieceAttackingAnotherPiece from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("f3", ""), []string{"e1", "e2", "f2", "f4", "h1", "h3"}) {
		t.Fatalf("Attackers for f3 != {e1, e2, f2, f4, h1, h3} in position(%s)", positionFen)
	}
}

func TestAttackersEveryPieceDefendingAnotherPiece(t *testing.T) {
	const positionFen = "2r5/1b1p4/1kp1q3/4n3/8/8/8/4K3 b - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersEveryPieceDefendingAnotherPiece from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("c6", ""), []string{"b6", "b7", "c8", "d7", "e5", "e6"}) {
		t.Fatalf("Attackers for c6 != {b6, b7, c8, d7, e5, e6} in position(%s)", positionFen)
	}
}

func TestAttackersEveryPieceDefendingEmptySquare(t *testing.T) {
	const positionFen = "B3k3/8/8/2K4R/3QPN2/8/8/8 w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersEveryPieceDefendingEmptySquare from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("d5", constants.ColorWhite), []string{"a8", "c5", "d4", "e4", "f4", "h5"}) {
		t.Fatalf("Attackers for d5 != {a8, c5, d4, e4, f4, h5} in position(%s)", positionFen)
	}
}

func TestAttackersPinnedPiecesStillAttackAndDefend(t *testing.T) {
	// Knight on c3 is pinned, but it's still attacking d4 and defending e5
	const positionFen = "r1bqkbnr/ppp2ppp/2np4/1B2p3/3PP3/5N2/PPP2PPP/RNBQK2R b KQkq - 0 4"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersPinnedPiecesStillAttackAndDefend from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("d4", constants.ColorBlack), []string{"c6", "e5"}) {
		t.Fatalf("Attackers for d4 != {c6, d5} in position(%s)", positionFen)
	}

	if !IncludesSameMember(game.Attackers("e5", constants.ColorBlack), []string{"c6", "d6"}) {
		t.Fatalf("Attackers for e5 != {c6, d6} in position(%s)", positionFen)
	}
}

func TestAttackersKingCanAttackDefendedPieces(t *testing.T) {
	const positionFen = "3k4/8/8/8/3b4/3R4/4Pq2/4K3 w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersKingCanAttackDefendedPieces from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("f2", constants.ColorWhite), []string{"e1"}) {
		t.Fatalf("Attackers for f2 != {e1} in position(%s)", positionFen)
	}
}

func TestAttackersLotOfAttackers(t *testing.T) {
	const positionFen = "5k2/8/3N1N2/2NBQQN1/3R1R2/2NPRPN1/3N1N2/4K3 w - - 0 1"
	game, err := chess.NewGameFromFen(positionFen)

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersLotOfAttackers from position %s. %v", positionFen, err)
	}

	if !IncludesSameMember(
		game.Attackers("e4", constants.ColorWhite),
		[]string{
			"c3", "c5", "d2", "d3", "d4", "d5", "d6",
			"e3", "e5", "f2", "f3", "f4", "f5", "f6",
			"g3", "g5",
		}) {
		t.Fatalf("Attackers for e4 != {c3, c5, d2, d3, d4, d5, d6, e3, e5, f2, f3, f4, f5, f6, g3, g5} in position(%s)", positionFen)
	}
}

func TestAttackersNoAttackers(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersNoAttackers. %v", err)
	}

	if len(game.Attackers("e4", constants.ColorWhite)) != 0 {
		t.Fatalf("Attackers for e4 != {}")
	}
}

func TestAttackersReadmeTests(t *testing.T) {
	game, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize game for test TestAttackersReadmeTests. %v", err)
	}

	if !IncludesSameMember(game.Attackers("f3", ""), []string{"e2", "g2", "g1"}) {
		t.Fatalf("Attackers for f3 != {e2, g2, g1}.")
	}

	if !IncludesSameMember(game.Attackers("e2", ""), []string{"d1", "e1", "f1", "g1"}) {
		t.Fatalf("Attackers for e2 != {d1, e1, f1, g1}.")
	}

	if len(game.Attackers("f6", "")) != 0 {
		t.Fatalf("Attackers for f6 != {}")
	}

	if _, err = game.Move("e4", false); err != nil {
		t.Fatalf("Failed to execute move `e4` in game for test TestAttackersReadmeTests. %v", err)
	}

	if !IncludesSameMember(game.Attackers("f6", ""), []string{"g8", "e7", "g7"}) {
		t.Fatalf("Attackers for f6 != {g8, e7, g7}.")
	}

	if !IncludesSameMember(game.Attackers("f3", constants.ColorWhite), []string{"g2", "d1", "g1"}) {
		t.Fatalf("Attackers for f3 != {g2, d1, g1}.")
	}

	const positionFen = "4k3/4n3/8/8/8/8/4R3/4K3 w - - 0 1"

	if err := game.Load(positionFen, false, false); err != nil {
		t.Fatalf("Failed to load position(%s) in game for test TestAttackersReadmeTests. %v", positionFen, err)
	}

	if !IncludesSameMember(game.Attackers("c6", constants.ColorBlack), []string{"e7"}) {
		t.Fatalf("Attackers for c6 != {e7}.")
	}
}
