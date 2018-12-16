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

// ToString ...
func (player *Player) ToString() string {
	return fmt.Sprintf(
		"name:%s\tpocket:%s\t%s\t%d",
		player.Name,
		player.Pocket.ToString(),
		player.BestHand().ToString(),
		"") //BestPossible(player.Pocket, player.Game.Board, Cards(player.Game.Deck))
}

// Players type
type Players []*Player
