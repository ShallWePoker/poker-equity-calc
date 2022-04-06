package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
	"time"
)

func main() {
	var removedCards []models.Card
	n := 10
	models.InitCardsEnums()
	fmt.Printf("cards to remove: %v\n", removedCards)
	fmt.Printf("generate %d hands\n", n)
	start := time.Now()
	for i := 0; i < n; i++ {
		_, err := utils.GenerateRandomNCards(7, removedCards)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("cards: %s\n", models.Hand(cards).ToString())

		// biggestMadeHand := utils.Seven2fiveV3(cards)
		// fmt.Printf("the biggest hand is %s\n", biggestMadeHand.BoardCards().ToString())
	}
	cost := time.Since(start)
	fmt.Println("1 cost:", cost)

	start2 := time.Now()
	for i := 0; i < n; i++ {
		// cards, err := utils.GenerateRandomNCards(7, cardsStr)
		_, err := utils.GenerateRandomNCardsV2(7, removedCards)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("cards: %s\n", models.Hand(cards).ToString())

		// biggestMadeHand := utils.Seven2fiveV3(cards)
		// fmt.Printf("the biggest hand is %s\n", biggestMadeHand.BoardCards().ToString())
	}
	cost2 := time.Since(start2)
	fmt.Println("2 cost:", cost2)
}
