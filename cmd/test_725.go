package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
	"strings"
)

func main() {
	var cards = flag.String("cards", "", "cards")
	flag.Parse()
	cardsStr := strings.Split(*cards, ",")
	tCards := make([]models.Card, 0)
	for _, c := range cardsStr {
		card, err := models.InitCardFromString(c)
		if err != nil {
			panic(err)
		}
		tCards = append(tCards, card)
	}
	hand := models.Hand{}
	for _, card := range tCards {
		hand = append(hand, card)

	}
	fmt.Printf("cards: %s\n", hand.ToString())
	fmt.Printf("[Seven2fiveV1] the biggest hand is %s\n", utils.Seven2five(tCards).BoardCards().ToString())
	fmt.Printf("[Seven2fiveV2] the biggest hand is %s\n", utils.Seven2fiveV2(tCards).BoardCards().ToString())
}
