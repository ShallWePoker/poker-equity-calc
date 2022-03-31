package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func main() {
	cards1 := []models.Card{
		{
			Title: "A",
			Suit:  "❤",
			Rank:  14,
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
			Title: "K",
			Suit:  "♣",
			Rank:  13,
		},
		{
			Title: "K",
			Suit:  "♠",
			Rank:  13,
		},
	}
	// cards2 := []models.Card{
	// 	{
	// 		Title: "A",
	// 		Suit:  "❤",
	// 		Rank:  14,
	// 	},
	// 	{
	// 		Title: "2",
	// 		Suit:  "♦",
	// 		Rank:  2,
	// 	},
	// 	{
	// 		Title: "5",
	// 		Suit:  "♣",
	// 		Rank:  5,
	// 	},
	// 	{
	// 		Title: "2",
	// 		Suit:  "♣",
	// 		Rank:  2,
	// 	},
	// 	{
	// 		Title: "K",
	// 		Suit:  "♣",
	// 		Rank:  13,
	// 	},
	// 	{
	// 		Title: "K",
	// 		Suit:  "♠",
	// 		Rank:  13,
	// 	},
	// 	{
	// 		Title: "8",
	// 		Suit:  "♣",
	// 		Rank:  8,
	// 	},
	// }
	hand := models.Hand{}
	for _, card := range cards1 {
		hand = append(hand, card)
	}
	fmt.Printf("cards: %s\n", hand.ToString())
	fmt.Printf("[Seven2fiveV1] the biggest hand is %s\n", utils.Seven2five(cards1).BoardCards().ToString())
	fmt.Printf("[Seven2fiveV2] the biggest hand is %s\n", utils.Seven2fiveV2(cards1).BoardCards().ToString())
}
