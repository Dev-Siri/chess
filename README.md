# Chess

Library for analysing and playing chess games programmatically. Direct translation of [chess.js](https://github.com/jhlywa/chess.js) to Go.

Represents a game with the 0x88 chess board representation.

## Why?

I know there's a library called `notnil/chess` for Go too, but I found one function not available in that package, the `.Moves()` function (That is available in chess.js) to count for mobility for one side.

So I directly translated the Chess.js library to Go. Also provides function for validating FENs and prettier move representation.

## holy hell

En passant still doesn't work properly (look at test "TestPositionsPgn") The captured pawn is still not taken off the board. (Wrong squared set to nil & I tried for a while but I don't know how to fix it.)

## Usage

- Get the package

```sh
go get github.com/Dev-Siri/chess
```

- Start using it ¯\\\_(ツ)\_/¯

```go
package main

import (
  "log"
  "fmt"

  "github.com/Dev-Siri/chess"
)

func main() {
  game, err := chess.NewGame()

  if err != nil {
    log.Fatalf(err)
  }

  game2, err := chess.NewGameFromFen("Or start from a FEN position")

  if err != nil {
    log.Fatalf(err)
  }

  // Make a move
  move, _ := game.Move("e4")

  // Print position
  fmt.Println(move.Pgn())
  fmt.Println(move.Fen())
}
```

## License

This project is MIT Licensed, see [LICENSE](LICENSE)
