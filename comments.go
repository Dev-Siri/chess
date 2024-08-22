package chess

import (
	"strings"

	"github.com/Dev-Siri/chess/schemas"
)

func (c *Chess) pruneComments() {
	reversedHistory := []*schemas.InternalMove{}
	currentComments := map[string]string{}

	copyComment := func(fen string) {
		if _, exists := currentComments[fen]; exists {
			currentComments[fen] = c.comments[fen]
		}
	}

	for len(c.history) > 0 {
		reversedHistory = append(reversedHistory, c.undoMove())
	}

	copyComment(c.Fen())

	for {
		move := reversedHistory[len(reversedHistory)-1]
		reversedHistory = reversedHistory[:len(reversedHistory)-1]

		if move == nil {
			break
		}

		c.makeMove(move)
		copyComment(c.Fen())
	}

	c.comments = currentComments
}

func (c *Chess) GetComment() string {
	return c.comments[c.Fen()]
}

func (c *Chess) SetComment(comment string) {
	c.comments[c.Fen()] = strings.ReplaceAll(strings.ReplaceAll(comment, "{", "]"), "}", "]")
}

func (c *Chess) GetComments() []schemas.Comment {
	c.pruneComments()

	comments := []schemas.Comment{}

	for fen, comment := range c.comments {
		comments = append(comments, schemas.Comment{
			Fen:     fen,
			Comment: comment,
		})
	}

	return comments
}

func (c *Chess) DeleteComments() []schemas.Comment {
	c.pruneComments()

	comments := []schemas.Comment{}

	for fen, comment := range c.comments {
		delete(c.comments, fen)
		comments = append(comments, schemas.Comment{
			Fen:     fen,
			Comment: comment,
		})
	}

	return comments
}
