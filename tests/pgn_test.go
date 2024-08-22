package tests

import (
	"os"
	"testing"

	"github.com/Dev-Siri/chess"
)

type Position struct {
	Moves         string
	Header        map[string]string
	MaxWidth      int
	Pgn           string
	NewlineChar   string
	StartPosition string
	Fen           string
}

var positions = []Position{
	{
		Moves: `d4 d5 Nf3 Nc6 e3 e6 Bb5 g5 O-O Qf6 Nc3 Bd7 Bxc6 Bxc6 Re1 O-O-O a4
		Bb4 a5 b5 axb6 axb6 Ra8+ Kd7 Ne5+ Kd6 Rxd8+ Qxd8 Nxf7+ Ke7 Nxd5+ Qxd5
		c3 Kxf7 Qf3+ Qxf3 gxf3 Bxf3 cxb4 e5 dxe5 Ke6 b3 Kxe5 Bb2+ Ke4 Bxh8 Nf6
		Bxf6 h5 Bxg5 Bg2 Kxg2 Kf5 Bh4 Kg4 Bg3 Kf5 e4+ Kg4 e5 h4 Bxh4 Kxh4 e6 c5
		bxc5 bxc5 e7 c4 bxc4 Kg4 e8=Q Kf5 Qe5+ Kg4 Re4#`,
		Header: map[string]string{
			"White":                   "Jeff Hlywa",
			"Black":                   "Steve Bragg",
			"GreatestGameEverPlayed?": "True",
		},
		MaxWidth:    19,
		NewlineChar: "<br />",
		Pgn:         fileToString("pgn/0.pgn"),
		Fen:         "8/8/8/4Q3/2P1R1k1/8/5PKP/8 b - - 4 39",
	},
	{
		Moves: `c4 e6 Nf3 d5 d4 Nf6 Nc3 Be7 Bg5 O-O e3 h6 Bh4 b6 cxd5 Nxd5 Bxe7
			Qxe7 Nxd5 exd5 Rc1 Be6 Qa4 c5 Qa3 Rc8 Bb5 a6 dxc5 bxc5 O-O Ra7 Be2 Nd7
			Nd4 Qf8 Nxe6 fxe6 e4 d4 f4 Qe7 e5 Rb8 Bc4 Kh8 Qh3 Nf8 b3 a5 f5 exf5
			Rxf5 Nh7 Rcf1 Qd8 Qg3 Re7 h4 Rbb7 e6 Rbc7 Qe5 Qe8 a4 Qd8 R1f2 Qe8 R2f3
			Qd8 Bd3 Qe8 Qe4 Nf6 Rxf6 gxf6 Rxf6 Kg8 Bc4 Kh8 Qf4`,
		Header: map[string]string{
			"Event":     "Reykjavik WCh",
			"Site":      "Reykjavik WCh",
			"Date":      "1972.01.07",
			"EventDate": "?",
			"Round":     "6",
			"Result":    "1-0",
			"White":     "Robert James Fischer",
			"Black":     "Boris Spassky",
			"ECO":       "D59",
			"WhiteElo":  "?",
			"BlackElo":  "?",
			"PlyCount":  "81",
		},
		MaxWidth: 65,
		Pgn:      fileToString("pgn/1.pgn"),
		Fen:      "4q2k/2r1r3/4PR1p/p1p5/P1Bp1Q1P/1P6/6P1/6K1 b - - 4 41",
	},
	{
		// testing MaxWidth being small and having no comments
		Moves:    `f3 e5 g4 Qh4#`,
		Header:   map[string]string{},
		MaxWidth: 1,
		Pgn:      fileToString("pgn/2.pgn"),
		Fen:      "rnb1kbnr/pppp1ppp/8/4p3/6Pq/5P2/PPPPP2P/RNBQKBNR w KQkq - 1 3",
	},
	{
		// testing a non-starting position
		Moves:         `Ba5 O-O d6 d4`,
		Header:        map[string]string{},
		MaxWidth:      20,
		Pgn:           fileToString("pgn/3.pgn"),
		StartPosition: "r1bqk1nr/pppp1ppp/2n5/4p3/1bB1P3/2P2N2/P2P1PPP/RNBQK2R b KQkq - 0 1",
		Fen:           "r1bqk1nr/ppp2ppp/2np4/b3p3/2BPP3/2P2N2/P4PPP/RNBQ1RK1 b kq - 0 3",
	},
}

func fileToString(path string) string {
	file, err := os.ReadFile("./" + path)

	if err != nil {
		return ""
	}

	return string(file)
}

func headersEqual(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TestPgnRemoveHeader(t *testing.T) {
	mainGame, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize comparisonGame for test TestPgnRemoveHeader. %v", err)
		return
	}

	const mainPgn = `
  [White "Paul Morphy"]
  [Black "Duke Karl / Count Isouard"]
  [fEn "1n2kb1r/p4ppp/4q3/4p1B1/4P3/8/PPP2PPP/2KR4 w k - 0 17"]

  17.Rd8# 1-0`

	comparisonGame, err := chess.NewGame()

	if err != nil {
		t.Fatalf("Failed to initialize comparison game for test TestPgnRemoveHeader. %v", err)
		return
	}

	const comparisonPgn = `
  [Black "Duke Karl / Count Isouard"]
  [fEn "1n2kb1r/p4ppp/4q3/4p1B1/4P3/8/PPP2PPP/2KR4 w k - 0 17"]

  17.Rd8# 1-0`

	mainGame.LoadPgn(mainPgn, false, "")
	comparisonGame.LoadPgn(comparisonPgn, false, "")

	if headersEqual(mainGame.Header(map[string]string{}), comparisonGame.Header(map[string]string{})) {
		t.Fatalf("mainGame.Header(map[string]string{}) == comparisonGame.Header(map[string]string{})). Header was not removed. Expected inequality.")
	}
}

func TestPositionsPgn(t *testing.T) {
	for i, position := range positions {
		t.Logf("positions[%d]", i)

		game, err := chess.NewGame()

		if err != nil {
			t.Fatalf("Failed to initialize chess.NewGame() for position %d", i)
			return
		}

		if position.StartPosition != "" {
			game.Load(position.StartPosition, false, false)
		}

		for i, move := range Split(position.Moves) {
			_, err := game.Move(move, false)

			if err != nil {
				t.Fatalf("Failed to execute move `%s` in game for position %d", move, i)
			}
		}

		if game.Fen() != position.Fen {
			t.Fatalf("Game FEN for position %d != positions[%d].Fen", i, i)
		}

		game.Header(position.Header)

		pgn := game.Pgn(position.NewlineChar, position.MaxWidth)

		if pgn != position.Pgn {
			t.Fatalf("Game PGN for position %d != positions[%d].Pgn", i, i)
		}
	}
}
