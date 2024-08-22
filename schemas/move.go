package schemas

type Move struct {
	Color string
	From  string
	To    string
	Piece string
	// nullable (can be "")
	Captured string
	// nullable (can be "")
	Promotion string
	Flags     string
	San       string
	Lan       string
	Before    string
	After     string
}
