package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func main() {
	cards1 := []models.Card{
		{
			Title: "T",
			Suit:  "❤",
			Rank:  10,
		},
		{
			Title: "Q",
			Suit:  "♠",
			Rank:  12,
		},
		{
			Title: "3",
			Suit:  "♦",
			Rank:  3,
		},
		{
			Title: "2",
			Suit:  "♦",
			Rank:  2,
		},
		{
			Title: "2",
			Suit:  "♣",
			Rank:  2,
		},
	}
	hand1, err := models.InitHand(cards1)
	if err != nil {
		panic(err)
	}
	fmt.Println(hand1.ToString())
	madeHand1 := hand1.Categorize()
	fmt.Printf("hand is : %s\n", madeHand1.Category())

	cards2 := []models.Card{
		{
			Title: "T",
			Suit:  "❤",
			Rank:  10,
		},
		{
			Title: "3",
			Suit:  "♠",
			Rank:  3,
		},
		{
			Title: "J",
			Suit:  "♦",
			Rank:  11,
		},
		{
			Title: "2",
			Suit:  "♦",
			Rank:  2,
		},
		{
			Title: "2",
			Suit:  "♣",
			Rank:  2,
		},
	}

	hand2, err := models.InitHand(cards2)
	if err != nil {
		panic(err)
	}
	fmt.Println(hand2.ToString())
	madeHand2 := hand2.Categorize()
	fmt.Printf("hand is : %s\n", madeHand2.Category())

	fmt.Printf("%s is greater than %s: %v\n", madeHand1.ToString(), madeHand2.ToString(), madeHand1.IsGreaterThan(madeHand2))
}
