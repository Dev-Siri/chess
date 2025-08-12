package constants

// Sides
const (
	ColorWhite = "w"
	ColorBlack = "b"
)

// Pieces
const (
	PiecePawn   = "p"
	PieceKnight = "n"
	PieceBishop = "b"
	PieceRook   = "r"
	PieceQueen  = "q"
	PieceKing   = "k"
)

const PieceSymbols = "pnbrqkPNBRQK"

// Terminations
const (
	TerminationWhiteVictory = "1-0"
	TerminationBlackVictory = "0-1"
	TerminationDraw         = "1/2-1/2"
	TerminationInProgress   = "*"
)

var TerminationMarkers = []string{TerminationWhiteVictory, TerminationBlackVictory, TerminationDraw, TerminationInProgress}
