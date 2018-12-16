package lib

// String ...
type String string

// ToCards ...
func (str String) ToCards() Cards {
	cards := make([]*Card, len(str)/2)
	for i := range cards {
		cards[i] = (&Card{}).FromString(string(str)[i*2 : i*2+2])
	}
	return Cards(cards)
}

// Bytes ...
type Bytes []byte

// ToCards ...
func (bytes Bytes) ToCards() Cards {
	cards := make([]*Card, len(bytes))
	for i := range bytes {
		cards[i] = (&Card{}).FromByte(bytes[i])
	}
	return Cards(cards)
}
