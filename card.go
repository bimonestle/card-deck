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
	Joker
)

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
