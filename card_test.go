package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Rank: Four, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Jack of Diamonds
	// Four of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}

}
