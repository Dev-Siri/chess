package chess

import (
	"fmt"

	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) Move(move string, strict bool) (*schemas.Move, error) {
	moveStruct := c.moveFromSan(move, strict)

	if moveStruct == nil {
		return nil, fmt.Errorf("invalid move: %s", move)
	}

	prettyMove := c.makePretty(moveStruct)

	c.makeMove(moveStruct)
	c.incPositionCount(prettyMove.After)

	return &prettyMove, nil
}
