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
	fmt.Printf("block cards: ")
	for _, blockCard := range blockCards {
		fmt.Printf("%s ", blockCard.ToString())
	}
	var multiRange []models.HoleCardRange
	for _, rge := range ranges {
		suited := false
		if len(rge) == 3 && string(rge[2]) == "s"{
			suited = true
		}
		holecardRange, err := models.InitHoleCardRange(rge[:2], suited)
		if err != nil {
			panic(err)
		}
		multiRange = append(multiRange, holecardRange)
	}
	fmt.Printf("\nrange: ")
	for _, rge := range multiRange {
		fmt.Printf("%s ", rge.ToString())
	}
	holeCards := utils.GenerateAllHoleCardFromMultiRange(multiRange, blockCards...)
	fmt.Print("\ngenerated hole card:\n")
	for i, holeCard := range holeCards {
		fmt.Printf("holecard #%d: %+v\n", i+1, holeCard.ToString())
	}
	fmt.Printf("total: %d\n", len(holeCards))
	endTime := time.Now().UnixMilli()
	fmt.Printf("time spent: %dms\n", endTime-startTime)
}

