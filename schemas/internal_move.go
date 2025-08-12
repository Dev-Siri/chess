package schemas

type InternalMove struct {
	Color string
	From  int
	To    int
	Piece string
	// nullable (can be "")
	Captured string
	// nullable (can be "")
	Promotion string
	Flags     int
}
