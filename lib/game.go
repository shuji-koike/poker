package lib

import (
	"fmt"
)

// Game ...
type Game struct {
	Deck    Deck
	Board   Cards
	Players Players
	Seed    int64
}

// Join ...
func (game *Game) Join(player *Player) {
	game.Players = append(game.Players, player)
}

// Deal ...
func (game *Game) Deal(cards *Cards) *Card {
	card := game.Deck[0]
	game.Deck = game.Deck[1:]
	*cards = append(*cards, card)
	return card
}

// Init ...
func (game *Game) Init() {
	game.Deck = make(Deck, DeckSize)
	game.Deck.Init()
	game.Deck.Shuffle(game.Seed)
	for i := range game.Players {
		game.Players[i].Game = game
	}
	for j := 0; j < PocketSize; j++ {
		for i := range game.Players {
			game.Deal(&game.Players[i].Pocket)
		}
	}
}

// Step ...
func (game *Game) Step() {
	if len(game.Board) == 0 {
		for i := 0; i < FlopSize; i++ {
			game.Deal(&game.Board)
		}
	} else if len(game.Board) < BoardSize {
		game.Deal(&game.Board)
	}
}

// String ...
func (game *Game) String() string {
	str := "----\n"
	for i := range game.Players {
		str = str + game.Players[i].String() + "\n"
	}
	str = str + fmt.Sprintf("board:%s\n", game.Board.String())
	str = str + fmt.Sprintf("deck:%s\n", Cards(game.Deck).String())
	str = str + fmt.Sprintf("seed:%d\n", game.Seed)
	return str
}
