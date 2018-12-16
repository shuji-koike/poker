package lib

import (
	"fmt"
	"sort"
)

// SortedCards ...
type SortedCards Cards

// Rank ...
func (cards Cards) Rank() Hand {
	hand := Hand{}
	if len(cards) != 5 {
		return hand
	}
	hand.Cards = cards.Copy().Sort()
	hand.Pairs = hand.Cards.Pairs()
	hand.isFlush = hand.IsFlush()
	hand.isStraight = hand.IsStraight()
	hand.highCard = hand.Cards[0]
	if !hand.isFlush && !hand.isStraight && len(hand.Pairs[0]) == 1 {
		hand.rank = HighCard
	} else if !hand.isFlush && len(hand.Pairs[0]) == 2 && len(hand.Pairs[1]) == 2 {
		hand.rank = TwoPair
	} else if !hand.isFlush && len(hand.Pairs[0]) == 2 {
		hand.rank = Pair
	} else if len(hand.Pairs[0]) == 3 && len(hand.Pairs[1]) == 2 {
		hand.rank = FullHouse
	} else if !hand.isFlush && len(hand.Pairs[0]) == 3 {
		hand.rank = ThreeOfAKind
	} else if len(hand.Pairs[0]) == 4 {
		hand.rank = FourOfAKind
	} else if hand.isStraight && hand.isFlush && hand.Cards[0].Rank == ACE && hand.Cards[1].Rank == KING {
		hand.rank = RoyalFlush
	} else if hand.isStraight && hand.isFlush {
		hand.rank = StrightFlush
	} else if hand.isFlush {
		hand.rank = Flush
	} else if hand.isStraight {
		hand.rank = Stright
	}
	return hand
}

// IsFlush ...
func (hand Hand) IsFlush() bool {
	for i := range hand.Cards {
		if hand.Cards[0].Suit != hand.Cards[i].Suit {
			return false
		}
	}
	return true
}

// IsStraight ...
func (hand Hand) IsStraight() bool {
	for i := range hand.Cards {
		a, b := hand.Cards[0], hand.Cards[i]
		if i == 1 && a.Rank == ACE && b.Rank == KING {
			continue
		} else if a.Rank != b.Rank+Rank(i) {
			return false
		}
	}
	return true
}

// Kicker ...
func (hand Hand) Kicker() *Card {
	return hand.Cards[4] //TODO
}

// ToString ...
func (hand Hand) ToString() string {
	return fmt.Sprintf(
		"hand:%s\trank:%d\tname:%s",
		Cards(hand.Cards).ToString(),
		hand.rank,
		HandNames[hand.rank])
}

// Sort ...
func (cards Cards) Sort() SortedCards {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank != cards[j].Rank {
			if cards[i].Rank == ACE {
				return true
			} else if cards[j].Rank == ACE {
				return false
			}
			return cards[i].Rank > cards[j].Rank
		}
		return cards[i].Suit > cards[j].Suit
	})
	return SortedCards(cards)
}

// Pairs ...
func (cards SortedCards) Pairs() []Cards {
	pairs := make([]Cards, 0)
	for i := range cards {
		flag := true
		for j := range pairs {
			if pairs[j][0].Rank == cards[i].Rank {
				pairs[j] = append(pairs[j], cards[i])
				flag = false
				break
			}
		}
		if flag {
			pairs = append(pairs, Cards{cards[i]})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return len(pairs[i]) > len(pairs[j])
	})
	return pairs
}
