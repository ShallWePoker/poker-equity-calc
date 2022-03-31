package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
	"time"
)

func main() {
	var removedCards []string
	var tList [][]models.Card
	n := 150000
	for i := 0; i < n; i++ {
		cards, err := utils.GenerateRandomNCards(7, removedCards)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("cards: %s\n", models.Hand(cards).ToString())
		tList = append(tList, cards)
	}

	startTime1 := time.Now().UnixNano()

	for _, cards := range tList {
		// fmt.Printf("[Seven2fiveV1]the biggest hand is %s\n", utils.Seven2five(cards).BoardCards().ToString())
		utils.Seven2five(cards).BoardCards().ToString()
	}
	endTime1 := time.Now().UnixNano()

	fmt.Printf("[Seven2fiveV1] time spent calculating: %ds\n", endTime1-startTime1)

	startTime2 := time.Now().UnixNano()
	for _, cards := range tList {
		// fmt.Printf("[Seven2fiveV2] the biggest hand is %s\n", utils.Seven2fiveV2(cards).BoardCards().ToString())
		utils.Seven2fiveV2(cards).BoardCards().ToString()
	}
	endTime2 := time.Now().UnixNano()

	fmt.Printf("[Seven2fiveV2] time spent calculating: %ds\n", endTime2-startTime2)
}
