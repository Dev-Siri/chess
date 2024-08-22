package chess

import "github.com/Dev-Siri/chess/constants"

func (c *Chess) Reset() error {
	return c.Load(constants.DefaultPositionFen, false, false)
}
