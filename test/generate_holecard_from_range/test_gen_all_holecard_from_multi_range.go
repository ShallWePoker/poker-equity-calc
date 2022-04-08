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
	var rangeStrInput = flag.String("multi_range", "", "multi range")
	var blockCardsInput = flag.String("block_cards", "", "block cards")
	flag.Parse()
	startTime := time.Now().UnixMilli()
	ranges := strings.Split(*rangeStrInput, ",")
	blockCardsStr := strings.Split(*blockCardsInput, ",")
	blockCards := make([]models.Card, 0)
	for _, blockCard := range blockCardsStr {
		card, err := models.InitCardFromString(blockCard)
		if err != nil {
			panic(err)
		}
		blockCards = append(blockCards, card)
	}
	var multiRange []models.HoleCardRange
	for _, rge := range ranges {
		holecardRange, err := models.InitHoleCardRange(rge, false)
		if err != nil {
			panic(err)
		}
		multiRange = append(multiRange, holecardRange)
	}
	holeCards := utils.GenerateAllHoleCardFromMultiRange(multiRange, blockCards...)
	fmt.Print("generated hole card:\n")
	for i, holeCard := range holeCards {
		fmt.Printf("holecard #%d: %+v\n", i+1, holeCard.ToString())
	}
	fmt.Printf("total: %d\n", len(holeCards))
	endTime := time.Now().UnixMilli()
	fmt.Printf("time spent: %dms\n", endTime-startTime)
}

