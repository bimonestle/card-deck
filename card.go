// go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // This is a special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart} // an Array of Suit

type Rank uint8

const (
	// By putting the blank identifier, Ace will hold a value of 1
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// It represents a card.
type Card struct {
	// Suit. The type: Clubs, Spade, Hearts, Diamond
	Suit
	// Rank.
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// Create a new deck of cards or shuffle(?)
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	// for each suit
	for _, suit := range suits {
		// for each rank
		for rank := minRank; rank <= maxRank; rank++ {
			// add card{suit, rank} to cards
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// Sort the cards from the smallest
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Check whether the first card is less than second one
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// Get the absolute rank of a single card
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Figure out what the new order of the cards should be
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))

	// Random source of numbers based on time
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// it will have all the numbers between 0 & that particular number
	perm := r.Perm(len(cards))
	for i, j := range perm {
		// j is the value of the inside of a permutation
		ret[i] = cards[j]
	}
	return ret
}
