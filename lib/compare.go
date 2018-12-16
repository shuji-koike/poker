package lib

// Compare ...
func (hand *Hand) Compare(a *Hand) Compare {
	if hand.rank > a.rank {
		return Win
	} else if hand.rank < a.rank {
		return Lose
	}
	return Tie
}
