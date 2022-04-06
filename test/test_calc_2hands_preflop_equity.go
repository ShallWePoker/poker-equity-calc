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
	var sampleRatio = flag.Int("ratio", 0, "inspect sample every ratio hands")
	var player1Hand = flag.String("player1", "", "player 1's hand")
	var player2Hand = flag.String("player2", "", "player 2's hand")
	flag.Parse()
	startTime := time.Now().Unix()
	fmt.Printf("generate %d sample boards\n", *n)
	fmt.Printf("inspect sample every %d hands\n", *sampleRatio)
	models.InitCardsEnums()
	player1HandStr := strings.Split(*player1Hand, ",")
	player2HandStr := strings.Split(*player2Hand, ",")

	player1HandCard1, err := models.InitCardFromString(player1HandStr[0])
	if err != nil {
		panic(err)
	}
	player1HandCard2, err := models.InitCardFromString(player1HandStr[1])
	if err != nil {
		panic(err)
	}
	player1HoleCards := make([]models.Card, 0)
	player1HoleCards = append(player1HoleCards, player1HandCard1, player1HandCard2)

	player2HandCard1, err := models.InitCardFromString(player2HandStr[0])
	if err != nil {
		panic(err)
	}
	player2HandCard2, err := models.InitCardFromString(player2HandStr[1])
	if err != nil {
		panic(err)
	}
	player2HoleCards := make([]models.Card, 0)
	player2HoleCards = append(player2HoleCards, player2HandCard1, player2HandCard2)


	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, player1HoleCards...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, player2HoleCards...)

	fmt.Printf("cardsToRemoveFromDeck: %v\n", cardsToRemoveFromDeck)

	player1Win := 0
	player2Win := 0
	tie := 0
	for i := 0; i < *n; i++ {
		time.Sleep(time.Nanosecond*42)
		generated5Cards, err := utils.GenerateRandomNCards(5, cardsToRemoveFromDeck)
		if err != nil {
			panic(err)
		}
		if i%*sampleRatio == 0 {
			fmt.Printf("board is: %s\n", models.Hand(generated5Cards).ToString())
		}
		player1BiggestHand := utils.Seven2five(append(generated5Cards, player1HoleCards...))
		player2BiggestHand := utils.Seven2five(append(generated5Cards, player2HoleCards...))
		if player1BiggestHand.IsGreaterThan(player2BiggestHand) {
			player1Win += 1
			if i%*sampleRatio == 0 {
				fmt.Printf("player1's %s wins player2's %s\n\n", player1BiggestHand.ToString(), player2BiggestHand.ToString())
			}
		} else if player2BiggestHand.IsGreaterThan(player1BiggestHand) {
			player2Win += 1
			if i%*sampleRatio == 0 {
				fmt.Printf("player1's %s loses player2's %s\n\n", player1BiggestHand.ToString(), player2BiggestHand.ToString())
			}
		} else {
			tie += 1
			if i%*sampleRatio == 0 {
				fmt.Printf("player1's%sties player2's %s\n\n", player1BiggestHand.ToString(), player2BiggestHand.ToString())
			}
		}
	}
	player1Equity := float64(player1Win) / float64(*n)
	player2Equity := float64(player2Win) / float64(*n)
	tieRate := float64(tie) / float64(*n)

	endTime := time.Now().Unix()
	fmt.Printf("%v equity: %v\n", models.HoleCards(player1HoleCards).ToString(), player1Equity)
	fmt.Printf("%v equity: %v\n", models.HoleCards(player2HoleCards).ToString(), player2Equity)
	fmt.Printf("tie rate: %v\n", tieRate)
	fmt.Printf("time spent calculating: %ds\n", endTime-startTime)
}
