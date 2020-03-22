// go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
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
func New() []Card {
	var cards []Card
	// for each suit
	for _, suit := range suits {
		// for each rank
		for rank := minRank; rank <= maxRank; rank++ {
			// add card{suit, rank} to cards
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}
