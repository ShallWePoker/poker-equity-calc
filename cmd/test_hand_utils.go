package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strings"
)

func main() {
	var inputCards = flag.String("cards", "", "10 cards!")
	flag.Parse()
	cardsStr := strings.Split(*inputCards, ",")
	fmt.Println(cardsStr)
	cards1 := make([]models.Card, 0)
	cards2 := make([]models.Card, 0)
	for i, input := range cardsStr {
		if len(input) != 2 {
			panic(fmt.Sprintf("input 2 letters representing 1 card. invalid input: %s", input))
		}
		title := string(input[0])
		suit := string(input[1])
		card := models.Card{Title: title, Suit: suit}
		err := card.Format()
		if err != nil {
			panic(err)
		}
		if i < 5 {
			cards1 = append(cards1, card)
			continue
		}
		cards2 = append(cards2, card)
	}
	hand1, err := models.InitHand(cards1)
	if err != nil {
		panic(err)
	}
	fmt.Println(hand1.ToString())
	madeHand1 := hand1.Categorize()
	fmt.Printf("hand is : %s\n", madeHand1.Category())

	hand2, err := models.InitHand(cards2)
	if err != nil {
		panic(err)
	}
	fmt.Println(hand2.ToString())
	madeHand2 := hand2.Categorize()
	fmt.Printf("hand is : %s\n", madeHand2.Category())

	fmt.Printf("%s is greater than %s: %v\n", madeHand1.ToString(), madeHand2.ToString(), madeHand1.IsGreaterThan(madeHand2))
}
