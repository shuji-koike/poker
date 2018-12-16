package lib

import (
	"strings"
)

// Rank ...
type Rank int8

// ToString ...
func (rank Rank) ToString() string {
	return string(CardRanks[rank])
}

// Suit ...
type Suit int8

// ToString ...
func (suit Suit) ToString() string {
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

// ToString ...
func (card *Card) ToString() string {
	if card == nil {
		return "??"
	}
	return card.Rank.ToString() + card.Suit.ToString()
}

// Cards ...
type Cards []*Card

// Copy ...
func (cards *Cards) Copy() *Cards {
	clone := make(Cards, len(*cards))
	copy(clone, *cards)
	return &clone
}

// ToString ...
func (cards Cards) ToString() string {
	str := ""
	for i := range cards {
		str = str + (cards)[i].ToString()
	}
	return str
}
