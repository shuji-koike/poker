package lib

import (
	"math/rand"
)

// Deck array of cards
type Deck Cards

// Init deck
func (deck Deck) Init() {
	for i := range deck {
		deck[i] = (&Card{}).FromByte(byte(i))
	}
}

// Shuffle deck
func (deck Deck) Shuffle(seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}
