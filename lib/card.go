package lib

import (
	"strings"
)

// Rank ...
type Rank int8

// String ...
func (rank Rank) String() string {
	return string(CardRanks[rank])
}

// Suit ...
type Suit int8

// String ...
func (suit Suit) String() string {
	return string(CardSuits[suit])
}

// Card ...
type Card struct {
	Rank Rank
	Suit Suit
}

// FromByte ...
func (card *Card) FromByte(value byte) *Card {
	card.Rank = Rank(value % 13)
	card.Suit = Suit(value / 13)
	return card
}

// FromString ...
func (card *Card) FromString(str string) *Card {
	card.Rank = Rank(strings.Index(CardRanks, str[0:1]))
	card.Suit = Suit(strings.Index(CardSuits, str[1:2]))
	return card
}

// String ...
func (card *Card) String() string {
	if card == nil {
		return "??"
	}
	return card.Rank.String() + card.Suit.String()
}

// Cards ...
type Cards []*Card

// Copy ...
func (cards *Cards) Copy() *Cards {
	clone := make(Cards, len(*cards))
	copy(clone, *cards)
	return &clone
}

// String ...
func (cards Cards) String() string {
	str := ""
	for i := range cards {
		str = str + (cards)[i].String()
	}
	return str
}
