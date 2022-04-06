package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strings"
)

func main() {
	var player1Hand = flag.String("player1", "", "player 1's hand")
	flag.Parse()
	models.InitCardsEnums()
	player1HandStr := strings.Split(*player1Hand, ",")
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

	holeCards, err := models.InitHoleCards(player1HoleCards)
	if err != nil {
		panic(err)
	}
	fmt.Printf("player1 hole cards: %s\n", holeCards.ToString())
}
