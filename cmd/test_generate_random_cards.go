package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
	"strings"
)

func main() {
	var removedCards = flag.String("removedcards", "", "cards to remove")
	var n = flag.Int("n", 0, "generate n random boards")
	flag.Parse()
	cardsStr := strings.Split(*removedCards, ",")
	fmt.Printf("cards to remove: %v\n", cardsStr)
	fmt.Printf("generate %d hands\n", *n)
	for i := 0; i < *n; i++ {
		cards, err := utils.GenerateRandomNCards(7, cardsStr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("cards: %s\n", models.Hand(cards).ToString())
		fmt.Println()
		biggestMadeHand := utils.Seven2five(cards)
		fmt.Printf("the biggest hand is %s\n", biggestMadeHand.BoardCards().ToString())
		fmt.Println()
	}
}
