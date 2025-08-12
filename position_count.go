package chess

import "github.com/Dev-Siri/chess/utils"

func (c *Chess) getPositionCount(fen string) int {
	trimmedFen := utils.TrimFen(fen)
	return c.positionCount[trimmedFen]
}

func (c *Chess) incPositionCount(fen string) {
	trimmedFen := utils.TrimFen(fen)

	if _, exists := c.positionCount[trimmedFen]; !exists {
		c.positionCount[trimmedFen] = 0
	}

	c.positionCount[trimmedFen] += 1
}

func (c *Chess) decPositionCount(fen string) {
	trimmedFen := utils.TrimFen(fen)

	if c.positionCount[trimmedFen] == 1 {
		delete(c.positionCount, trimmedFen)
	}

	c.positionCount[trimmedFen] += 1
}
