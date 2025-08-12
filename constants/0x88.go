package constants

const (
	BitNormal          = "normal"
	BitCapture         = "capture"
	BitBigPawn         = "bigPawn"
	BitEpCapture       = "epCapture"
	BitPromotion       = "promotion"
	BitKingSideCastle  = "kingSideCastle"
	BitQueenSideCastle = "queenSideCastle"
)

var Bits = map[string]int{
	BitNormal:          1,
	BitCapture:         2,
	BitBigPawn:         4,
	BitEpCapture:       8,
	BitPromotion:       16,
	BitKingSideCastle:  32,
	BitQueenSideCastle: 64,
}

var Ox88 = map[string]int{
	"a8": 0, "b8": 1, "c8": 2, "d8": 3, "e8": 4, "f8": 5, "g8": 6, "h8": 7,
	"a7": 16, "b7": 17, "c7": 18, "d7": 19, "e7": 20, "f7": 21, "g7": 22, "h7": 23,
	"a6": 32, "b6": 33, "c6": 34, "d6": 35, "e6": 36, "f6": 37, "g6": 38, "h6": 39,
	"a5": 48, "b5": 49, "c5": 50, "d5": 51, "e5": 52, "f5": 53, "g5": 54, "h5": 55,
	"a4": 64, "b4": 65, "c4": 66, "d4": 67, "e4": 68, "f4": 69, "g4": 70, "h4": 71,
	"a3": 80, "b3": 81, "c3": 82, "d3": 83, "e3": 84, "f3": 85, "g3": 86, "h3": 87,
	"a2": 96, "b2": 97, "c2": 98, "d2": 99, "e2": 100, "f2": 101, "g2": 102, "h2": 103,
	"a1": 112, "b1": 113, "c1": 114, "d1": 115, "e1": 116, "f1": 117, "g1": 118, "h1": 119,
}

var PawnOffsets = map[string][4]int{
	ColorBlack: {16, 32, 17, 15},
	ColorWhite: {-16, -32, -17, -15},
}

var PieceOffsets = map[string][]int{
	PieceKnight: {-18, -33, -31, -14, 18, 33, 31, 14},
	PieceBishop: {-17, -15, 17, 15},
	PieceRook:   {-16, 1, 16, -1},
	PieceQueen:  {-17, -16, -15, 1, 17, 16, 15, -1},
	PieceKing:   {-17, -16, -15, 1, 17, 16, 15, -1},
}

var Attacks = []int{
	20, 0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 20, 0,
	0, 20, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 20, 0, 0,
	0, 0, 20, 0, 0, 0, 0, 24, 0, 0, 0, 0, 20, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 24, 0, 0, 0, 20, 0, 0, 0, 0,
	0, 0, 0, 0, 20, 0, 0, 24, 0, 0, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 20, 2, 24, 2, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 2, 53, 56, 53, 2, 0, 0, 0, 0, 0, 0,
	24, 24, 24, 24, 24, 24, 56, 0, 56, 24, 24, 24, 24, 24, 24, 0,
	0, 0, 0, 0, 0, 2, 53, 56, 53, 2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 20, 2, 24, 2, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 20, 0, 0, 24, 0, 0, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 24, 0, 0, 0, 20, 0, 0, 0, 0,
	0, 0, 20, 0, 0, 0, 0, 24, 0, 0, 0, 0, 20, 0, 0, 0,
	0, 20, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 20, 0, 0,
	20, 0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 20,
}

var Rays = []int{
	17, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 15, 0,
	0, 17, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 15, 0, 0,
	0, 0, 17, 0, 0, 0, 0, 16, 0, 0, 0, 0, 15, 0, 0, 0,
	0, 0, 0, 17, 0, 0, 0, 16, 0, 0, 0, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 0, 0, 16, 0, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 0, 16, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 17, 16, 15, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 0, -1, -1, -1, -1, -1, -1, -1, 0,
	0, 0, 0, 0, 0, 0, -15, -16, -17, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, -15, 0, -16, 0, -17, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, -15, 0, 0, -16, 0, 0, -17, 0, 0, 0, 0, 0,
	0, 0, 0, -15, 0, 0, 0, -16, 0, 0, 0, -17, 0, 0, 0, 0,
	0, 0, -15, 0, 0, 0, 0, -16, 0, 0, 0, 0, -17, 0, 0, 0,
	0, -15, 0, 0, 0, 0, 0, -16, 0, 0, 0, 0, 0, -17, 0, 0,
	-15, 0, 0, 0, 0, 0, 0, -16, 0, 0, 0, 0, 0, 0, -17,
}

var Promotions = []string{PieceKnight, PieceBishop, PieceRook, PieceQueen}

const (
	Rank1 = 7
	Rank2 = 6
	Rank7 = 1
	Rank8 = 0
)

var Sides = map[string]int{
	PieceKing:  Bits[BitKingSideCastle],
	PieceQueen: Bits[BitQueenSideCastle],
}

var Rooks = map[string][]map[string]int{
	ColorWhite: {
		{"square": Ox88["a1"], "flag": Bits[BitQueenSideCastle]},
		{"square": Ox88["h1"], "flag": Bits[BitKingSideCastle]},
	},
	ColorBlack: {
		{"square": Ox88["a8"], "flag": Bits[BitQueenSideCastle]},
		{"square": Ox88["h8"], "flag": Bits[BitKingSideCastle]},
	},
}

var SecondRank = map[string]int{
	ColorBlack: Rank7,
	ColorWhite: Rank2,
}

var PieceMasks = map[string]int{
	"p": 0x1,
	"n": 0x2,
	"b": 0x4,
	"r": 0x8,
	"q": 0x10,
	"k": 0x20,
}

var Flags = map[string]string{
	BitNormal:          "n",
	BitCapture:         "c",
	BitBigPawn:         "b",
	BitEpCapture:       "e",
	BitPromotion:       "p",
	BitKingSideCastle:  "k",
	BitQueenSideCastle: "q",
}
