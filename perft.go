package chess

func (c *Chess) Perft(depth int) int {
	moves := c.moves(false, "", "")
	var nodes int

	color := c.turn

	for i := 0; i < len(moves); i++ {
		c.makeMove(&moves[i])

		if !c.isKingAttacked(color) {
			if (depth - 1) > 0 {
				nodes += c.Perft(depth - 1)
			} else {
				nodes++
			}
		}

		c.undoMove()
	}

	return nodes
}
