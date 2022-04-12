package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/equity_calc"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strings"
	"time"
)

func main()  {
	var player1Hand = flag.String("player1", "", "player 1's hand")
	var player2rangeStrInput = flag.String("player2", "", "player 2's range")
	var flopCards = flag.String("flop", "", "flop cards")
	var turnCardStr = flag.String("turn", "", "turn card")
	var riverCardStr = flag.String("river", "", "river card")
	flag.Parse()
	startTime := time.Now().UnixMilli()
	models.InitCardsEnums()
	flopCardsStr := strings.Split(*flopCards, ",")
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
	fmt.Printf("player1 holecard: %s\n", models.HoleCard(player1HoleCards).ToString())
	player2ranges := strings.Split(*player2rangeStrInput, ",")
	var player2Rge []models.HoleCardRange
	for _, rge := range player2ranges {
		suited := false
		if len(rge) == 3 && string(rge[2]) == "s"{
			suited = true
		}
		holecardRange, err := models.InitHoleCardRange(rge[:2], suited)
		if err != nil {
			panic(err)
		}
		player2Rge = append(player2Rge, holecardRange)
	}
	fmt.Printf("\nplayer2 range: ")
	for _, rge := range player2Rge {
		fmt.Printf("%s ", rge.ToString())
	}
	flopCardsSlice := make([]models.Card, 0)
	for _, cardStr := range flopCardsStr {
		card, err := models.InitCardFromString(cardStr)
		if err != nil {
			panic(err)
		}
		flopCardsSlice = append(flopCardsSlice, card)
	}

	flop, err := models.InitFlop(flopCardsSlice)
	if err != nil {
		panic(err)
	}
	turnCard, err := models.InitCardFromString(*turnCardStr)
	if err != nil {
		panic(err)
	}
	riverCard, err := models.InitCardFromString(*riverCardStr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")
	player1Equity, player2Equity, err := equity_calc.HoleCardVersusBalancedRangeRiverEquity(player1HoleCards, player2Rge, flop, turnCard, riverCard)
	if err != nil {
		panic(err)
	}
	endTime := time.Now().UnixMilli()
	fmt.Printf("\nflop: %s turn: %s river: %s\n", flop.ToString(), turnCard.ToString(), riverCard.ToString())
	fmt.Printf("\nplayer1: %s\n", player1Equity.ToString())
	fmt.Printf("\nplayer2: %s\n", player2Equity.ToString())
	fmt.Printf("\ntime spent: %dms\n", endTime-startTime)
}




