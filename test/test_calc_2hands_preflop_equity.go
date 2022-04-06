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

	player1Equity, player2Equity, err := utils.HandVersusHandPreflopEquity(player1HoleCards, player2HoleCards)
	if err != nil {
		panic(err)
	}


	endTime := time.Now().Unix()
	fmt.Printf("%v equity: %v\n", models.HoleCards(player1HoleCards).ToString(), player1Equity.WinRate)
	fmt.Printf("%v equity: %v\n", models.HoleCards(player2HoleCards).ToString(), player2Equity.WinRate)
	fmt.Printf("tie rate: %v\n", player1Equity.TieRate)
	fmt.Printf("time spent calculating: %ds\n", endTime-startTime)
}
