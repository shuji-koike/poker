package lib

import (
	"sort"
)

// Hand ...
type Hand struct {
	Cards      SortedCards
	Pairs      []Cards
	highCard   *Card
	kicker     *Card
	isFlush    bool
	isStraight bool
	rank       int32
}

// Hands ...
func Hands(pocket, board Cards) []Hand {
	n := HandSize - PocketSize
	if len(board) < n {
		return nil
	}
	comb := board.Comb(n)
	for i := range comb {
		comb[i] = append(pocket, comb[i]...)
	}
	hands := make([]Hand, len(comb))
	for i := range comb {
		hands[i] = comb[i].Rank()
	}
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]
		return a.Compare(&b) == Win
	})
	// fmt.Println("//--")
	// for i := range hands {
	// 	Print(hands[i])
	// }
	return hands
}

// BestPossible ...
func BestPossible(pocket, board, deck Cards) []Hand {
	if len(pocket) < PocketSize || len(board) < FlopSize {
		return []Hand{}
	}
	comb := make([]Cards, 0)
	for i := 0; i < BoardSize-len(board)+1; i++ {
		a := board.Comb(HandSize - PocketSize - i)
		b := deck.Comb(i)
		println(len(a), len(b))
		for j := range a {
			for k := range b {
				comb = append(comb, append(append(pocket, a[j]...), b[k]...))
			}
		}
	}
	if len(comb) == 0 {
		return Hands(pocket, board)
	}
	hands := make([]Hand, len(comb))
	for i := range comb {
		hands[i] = comb[i].Rank()
	}
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]
		return a.Compare(&b) == Win
	})
	return hands
}

// Comb ...
func (cards *Cards) Comb(n int) []Cards {
	if n < 1 {
		return []Cards{}
	}
	m := len(*cards)
	patterns := make([]Cards, 0)
	pattern := make(Cards, n)
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < m; j++ {
			pattern[i] = (*cards)[j]
			if i == n-1 {
				patterns = append(patterns, *pattern.Copy())
			} else {
				rc(i+1, j+1)
			}
		}
	}
	rc(0, 0)
	return patterns
}
