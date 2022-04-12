package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/equity_calc"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strings"
	"time"
)

func main() {
	var player1Hand = flag.String("player1", "", "player 1's hand")
	var player2Hand = flag.String("player2", "", "player 2's hand")
	flag.Parse()
	startTime := time.Now().Unix()
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

	player1Equity, player2Equity, err := equity_calc.HandVersusHandPreflopEquity(player1HoleCards, player2HoleCards)
	if err != nil {
		panic(err)
	}


	endTime := time.Now().Unix()

	fmt.Printf("preflop\n")
	fmt.Printf("%s\n", player1Equity.ToString())
	fmt.Printf("%s\n", player2Equity.ToString())

	fmt.Printf("time spent calculating: %ds\n", endTime-startTime)
}
