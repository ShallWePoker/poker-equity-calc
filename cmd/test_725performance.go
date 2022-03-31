package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func main() {
	var removedCards []string
	n := 10
	for i := 0; i < n; i++ {
		cards, err := utils.GenerateRandomNCards(7, removedCards)
		if err != nil {
			panic(err)
		}
		fmt.Printf("cards: %s\n", models.Hand(cards).ToString())
		fmt.Printf("[Seven2fiveV1]the biggest hand is %s\n", utils.Seven2five(cards).BoardCards().ToString())
		fmt.Printf("[Seven2fiveV2]the biggest hand is %s\n", utils.Seven2fiveV2(cards).BoardCards().ToString())
	}
}
