package utils

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"unicode"

	"github.com/Dev-Siri/chess/constants"
)

func IsDigit(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func SwapColors(color string) string {
	if color == constants.ColorWhite {
		return constants.ColorBlack
	} else {
		return constants.ColorWhite
	}
}

func Some[T any](slice []T, test func(T) bool) bool {
	for _, v := range slice {
		if test(v) {
			return true
		}
	}
	return false
}

func ToHex(str string) string {
	var sb strings.Builder

	for _, c := range str {
		if c < 128 {
			sb.WriteString(fmt.Sprintf("%02x", c))
		} else {
			escaped := url.QueryEscape(string(c))
			sb.WriteString(strings.ReplaceAll(strings.ToLower(escaped), "%", ""))
		}
	}

	return sb.String()
}

func FromHex(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	var sb strings.Builder
	for i := 0; i < len(str); i += 2 {
		pair := str[i : i+2]
		decoded, err := hex.DecodeString(pair)
		if err != nil {
			return "", err
		}
		sb.Write(decoded)
	}

	decodedStr, err := url.QueryUnescape(sb.String())
	if err != nil {
		return "", err
	}

	return decodedStr, nil
}

func LetterToChessUnicode(letter string) string {
	switch letter {
	case "P":
		return "♙"
	case "N":
		return "♘"
	case "R":
		return "♖"
	case "B":
		return "♗"
	case "K":
		return "♔"
	case "Q":
		return "♕"
	case "p":
		return "♟"
	case "n":
		return "♞"
	case "r":
		return "♜"
	case "b":
		return "♝"
	case "k":
		return "♚"
	case "q":
		return "♛"
	default:
		return ""
	}
}
