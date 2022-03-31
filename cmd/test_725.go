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
			Rank:  3,
		},
		{
			Title: "4",
			Suit:  "♣",
			Rank:  4,
		},
	}
	hand := models.Hand{}
	for _, card := range cards1 {
		hand = append(hand, card)
	}
	fmt.Printf("cards: %s\n", hand.ToString())
	hand1 := utils.Seven2fiveV2(cards1)
	fmt.Printf("the biggest hand is %s\n", hand1.BoardCards().ToString())
}
