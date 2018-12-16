package lib

// DeckSize ...
const DeckSize = 52

// HandSize ...
const HandSize = 5

// PocketSize ...
const PocketSize = 2

// BoardSize ...
const BoardSize = 5

// FlopSize ...
const FlopSize = 3

// CardRanks ...
const CardRanks = "A23456789TJQK"

// CardSuits ...
const CardSuits = "shdc"

// Rank enum
const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

// Suit enum
const (
	SPADES Suit = iota
	HEARTS
	DIAMONDS
	CLUBS
)

// HandNames ...
var HandNames = []string{
	"HighCard",
	"Pair",
	"TwoPair",
	"ThreeOfAKind",
	"Stright",
	"Flush",
	"FullHouse",
	"FourOfAKind",
	"StrightFlush",
	"RoyalFlush"}

// Rank enum
const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	Stright
	Flush
	FullHouse
	FourOfAKind
	StrightFlush
	RoyalFlush
)

// Compare ...
type Compare int8

// Compare enum
const (
	Win  Compare = 1
	Tie          = 0
	Lose         = -1
)
