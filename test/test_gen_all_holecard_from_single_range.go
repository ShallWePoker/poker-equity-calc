package main

import (
	"flag"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func main() {
	var rangeStr = flag.String("range", "", "player 1's hand")
	flag.Parse()
	c1 := *rangeStr
	hcr, err := models.InitHoleCardRange(c1, false)
	if err != nil {
		panic(err)
	}
	holeCards := utils.GenerateAllHoleCardFromSingleRange(hcr)
	fmt.Print("generated hole card:\n")
	for i, holeCard := range holeCards {
		fmt.Printf("holecard #%d: %+v\n", i+1, holeCard.ToString())
	}
	fmt.Printf("total: %d\n", len(holeCards))
}
