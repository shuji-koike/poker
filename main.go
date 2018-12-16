package main

import (
	"fmt"

	"github.com/shuji-koike/poker/lib"
)

func main() {
	lib.Print(lib.String("AsKsQsJsTs").ToCards().Rank())
	lib.Print(lib.String("KsQsJsTs9s").ToCards().Rank())
	lib.Print(lib.String("AsAhAdAcTs").ToCards().Rank())
	lib.Print(lib.String("KsQsJsTs8s").ToCards().Rank())
	lib.Print(lib.String("KsQsJsTs9d").ToCards().Rank())
	lib.Print(lib.String("KsKdKhQsQd").ToCards().Rank())
	lib.Print(lib.String("KsKdKhQsJd").ToCards().Rank())
	lib.Print(lib.String("KsKdThQsQd").ToCards().Rank())
	lib.Print(lib.Bytes{0, 1, 2, 3, 4}.ToCards().Rank())

	game := lib.Game{Seed: 5}
	game.Join(&lib.Player{Name: "Alice"})
	game.Join(&lib.Player{Name: "Bob"})
	game.Join(&lib.Player{Name: "Charlie"})
	game.Join(&lib.Player{Name: "Dave"})
	game.Join(&lib.Player{Name: "Eve"})
	game.Init()
	game.Step()
	fmt.Print(game.ToString())
	game.Step()
	fmt.Print(game.ToString())
	game.Step()
	fmt.Print(game.ToString())
}
