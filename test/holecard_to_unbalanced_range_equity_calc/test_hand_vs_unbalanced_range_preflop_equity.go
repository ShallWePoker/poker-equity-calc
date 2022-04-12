package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/equity_calc"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strconv"
	"strings"
	"time"
)

func main()  {
	var player1Hand = flag.String("player1", "", "player 1's hand")
	// AA,AKs:0.7@22,A5s:0.3
	var player2rangeStrInput = flag.String("player2", "", "player 2's range")
	flag.Parse()
	fmt.Printf("\n%s\n", *player2rangeStrInput)
	startTime := time.Now().UnixMilli()
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
	fmt.Printf("player1 holecard: %s\n", models.HoleCard(player1HoleCards).ToString())
	var player2Rge []models.UnbalancedHolecardRange
	player2unbalancedRanges := strings.Split(*player2rangeStrInput, "@")
	for _, rge := range player2unbalancedRanges {
		ranges := make([]models.HoleCardRange, 0)
		unbalancedRange := models.UnbalancedHolecardRange{}
		rangeParts := strings.Split(rge, ":")
		percentageStr := rangeParts[1]
		percentage, err := strconv.ParseFloat(percentageStr, 64)
		if err != nil {
			panic(err)
		}
		unbalancedRange.Percentage = percentage
		rangePartStr := rangeParts[0]
		rangePart := strings.Split(rangePartStr, ",")
		for _, range1 := range rangePart {
			suited := false
			if len(range1) == 3 && string(range1[2]) == "s"{
				suited = true
			}
			holecardRange, err := models.InitHoleCardRange(range1[:2], suited)
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, holecardRange)
		}
		unbalancedRange.HolecardRanges = ranges
		player2Rge = append(player2Rge, unbalancedRange)
	}
	fmt.Printf("\nplayer2 range: %s\n", models.UnbalancedHolecardRanges(player2Rge).ToString())
	fmt.Printf("\n")
	player1Equity, player2Equity, err := equity_calc.HoleCardVersusUnbalancedRangePreflopEquity(player1HoleCards, player2Rge)
	if err != nil {
		panic(err)
	}
	endTime := time.Now().UnixMilli()
	fmt.Printf("\npreflop:\n")
	fmt.Printf("\nplayer1: %s\n", player1Equity.ToString())
	fmt.Printf("\nplayer2: %s\n", player2Equity.ToString())
	fmt.Printf("\ntime spent: %dms\n", endTime-startTime)
}


