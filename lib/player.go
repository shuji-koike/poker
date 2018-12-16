package lib

import "fmt"

// Player struct
type Player struct {
	Name   string
	Game   *Game
	Pocket Cards
}

// BestHand ...
func (player *Player) BestHand() Hand {
	return Hands(player.Pocket, player.Game.Board)[0]
}

// String ...
func (player *Player) String() string {
	return fmt.Sprintf(
		"name:%s\tpocket:%s\t%s\t%s",
		player.Name,
		player.Pocket.String(),
		player.BestHand().String(),
		BestPossible(player.Pocket, player.Game.Board, Cards(player.Game.Deck))[0].String())
}

// Players type
type Players []*Player
