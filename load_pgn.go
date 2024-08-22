package chess

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Dev-Siri/chess/constants"
	"github.com/Dev-Siri/chess/utils"
)

func mask(str string) string {
	return strings.ReplaceAll(str, "\\", "\\\\")
}

func encodeComment(str string, newlineChar string) string {
	str = strings.ReplaceAll(str, mask(newlineChar), " ")
	return "{" + utils.ToHex(str[1:len(str)-1]) + "}"
}

func decodeComment(str string) (string, error) {
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		hex, err := utils.FromHex(str[1 : len(str)-1])

		if err != nil {
			return "", err
		}

		return hex, nil
	}

	return "", nil
}

func parsePgnHeaders(header string, newlineChar string) map[string]string {
	headerMap := map[string]string{}
	headersRe := regexp.MustCompile(mask(newlineChar))
	headers := headersRe.Split(header, -1)

	key, value := "", ""

	for i := 0; i < len(headers); i++ {
		regex := regexp.MustCompile(`^\s*\[\s*([A-Za-z]+)\s*"(.*)"\s*\]\s*$`)
		key, value = regex.ReplaceAllString(headers[i], "$1"), regex.ReplaceAllString(headers[i], "$2")

		if len(strings.TrimSpace(key)) > 0 {
			headerMap[key] = value
		}
	}

	return headerMap
}

func (c *Chess) LoadPgn(pgn string, strict bool, newlineChar string) error {
	if newlineChar == "" {
		newlineChar = "\r?\n"
	}

	pgn = strings.TrimSpace(pgn)

	headerRegex := regexp.MustCompile(
		`^(\[((?:` +
			mask(newlineChar) +
			`)|.)*\])` +
			`((?:\s*` +
			mask(newlineChar) +
			`){2}|(?:\s*` +
			mask(newlineChar) +
			`)*$)`,
	)

	headerRegexResults := headerRegex.FindStringSubmatch(pgn)

	var headerString string

	if len(headerRegexResults) >= 2 {
		headerString = headerRegexResults[1]
	}

	c.Reset()

	headers := parsePgnHeaders(headerString, newlineChar)
	var fen string

	for key, value := range headers {
		if strings.ToLower(key) == "fen" {
			fen = headers[key]
		}

		c.Header(map[string]string{key: value})
	}

	if !strict {
		if fen != "" {
			c.Load(fen, false, true)
		}
	} else {
		if headers["SetUp"] == "1" {
			if _, exists := headers["FEN"]; !exists {
				return fmt.Errorf("invalid pgn: FEN tag must be supplied with SetUp tag")
			}

			c.Load(headers["FEN"], false, true)
		}
	}

	ms := strings.ReplaceAll(pgn, headerString, "")

	msRe := regexp.MustCompile(`({[^}]*})+?|;([^` + mask(newlineChar) + `]*)`)
	ms = msRe.ReplaceAllStringFunc(ms, func(match string) string {
		bracketMatch := regexp.MustCompile(`{[^}]*}`).FindString(match)
		if bracketMatch != "" {
			return encodeComment(bracketMatch, newlineChar)
		}
		semicolonMatch := match[1:]
		return " " + encodeComment("{"+semicolonMatch+"}", newlineChar)
	})

	ms = strings.ReplaceAll(ms, mask(newlineChar), " ")

	ravRe := regexp.MustCompile(`(\([^()]+\))+?`)

	for ravRe.MatchString(ms) {
		ms = ravRe.ReplaceAllString(ms, "")
	}

	ms = regexp.MustCompile(`\d+\.(\.\.)?`).ReplaceAllString(ms, "")

	ms = regexp.MustCompile(`\.\.\.`).ReplaceAllString(ms, "")

	ms = regexp.MustCompile(`\$\d+`).ReplaceAllString(ms, "")

	moves := strings.Fields(ms)
	filteredMoves := []string{}

	for _, move := range moves {
		if move != "" {
			filteredMoves = append(filteredMoves, move)
		}
	}

	moves = filteredMoves

	var result string

	for halfMove := 0; halfMove < len(moves); halfMove++ {
		comment, err := decodeComment(moves[halfMove])

		if err != nil {
			return err
		}

		if comment != "" {
			c.comments[c.Fen()] = comment
			continue
		}

		move := c.moveFromSan(moves[halfMove], strict)

		if move == nil {
			if slices.Contains(constants.TerminationMarkers, moves[halfMove]) {
				result = moves[halfMove]
			} else {
				return fmt.Errorf("invalid move in PGN: %s", moves[halfMove])
			}
		} else {
			result = ""
			c.makeMove(move)
			c.incPositionCount(c.Fen())
		}
	}

	_, resultExists := c.header["Result"]

	if result != "" && len(headers) > 0 && !resultExists {
		c.Header(map[string]string{"Result": result})
	}

	return nil
}
