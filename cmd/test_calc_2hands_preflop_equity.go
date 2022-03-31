package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
	"strings"
	"time"
)

func main() {
	var n = flag.Int("n", 0, "generate n random boards")
	var player1Hand = flag.String("player1", "", "player 1's hand")
	var player2Hand = flag.String("player2", "", "player 2's hand")
	flag.Parse()
	startTime := time.Now().Unix()
	fmt.Printf("generate %d sample boards\n", *n)
	player1HandStr := strings.Split(*player1Hand, ",")
	player2HandStr := strings.Split(*player2Hand, ",")

	cardsToRemoveFromDeck := append(player1HandStr, player2HandStr...)
	fmt.Printf("cardsToRemoveFromDeck: %v\n", cardsToRemoveFromDeck)

	player1Cards := make([]models.Card, 0)
	player2Cards := make([]models.Card, 0)
	for _, cardStr := range player1HandStr {
		card, err := models.InitCardFromString(cardStr)
		if err != nil {
			panic(err)
		}
		player1Cards = append(player1Cards, card)
	}

	for _, cardStr := range player2HandStr {
		card, err := models.InitCardFromString(cardStr)
		if err != nil {
			panic(err)
		}
		player2Cards = append(player2Cards, card)
	}
	fmt.Printf("player 1 hand: %v\n", player1Cards)
	fmt.Printf("player 2 hand: %v\n", player2Cards)

	player1Win := 0
	player2Win := 0
	tie := 0
	for i := 0; i < *n; i++ {
		generated5Cards, err := utils.GenerateRandomNCards(5, cardsToRemoveFromDeck)
		if err != nil {
			panic(err)
		}
		player1BiggestHand := utils.Seven2five(append(generated5Cards, player1Cards...))
		player2BiggestHand := utils.Seven2five(append(generated5Cards, player2Cards...))
		if player1BiggestHand.IsGreaterThan(player2BiggestHand) {
			player1Win += 1
		} else if player2BiggestHand.IsGreaterThan(player1BiggestHand) {
			player2Win += 1
		} else {
			tie += 1
		}
	}
	player1Equity := float64(player1Win) / float64(*n)
	player2Equity := float64(player2Win) / float64(*n)
	tieRate := float64(tie) / float64(*n)

	endTime := time.Now().Unix()
	fmt.Printf("%v equity: %v\n", player1HandStr, player1Equity)
	fmt.Printf("%v equity: %v\n", player2HandStr, player2Equity)
	fmt.Printf("tie rate: %v\n", tieRate)
	fmt.Printf("time spent calculating: %ds\n", endTime-startTime)
}
