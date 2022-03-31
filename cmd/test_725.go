package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
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
		{
			Title: "3",
			Suit:  "♣",
			Rank:  2,
		},
		{
			Title: "4",
			Suit:  "♣",
			Rank:  2,
		},
	}
	hand1 := utils.Seven2five(cards1)
	fmt.Println(hand1.ToString())
}
