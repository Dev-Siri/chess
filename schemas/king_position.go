package schemas

import "regexp"

type KingPosition struct {
	Color string
	Regex *regexp.Regexp
}
