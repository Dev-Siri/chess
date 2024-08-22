package schemas

type History struct {
	Move       *InternalMove
	Kings      map[string]int
	Turn       string
	Castling   map[string]int
	EpSquare   int
	HalfMoves  int
	MoveNumber int
}
