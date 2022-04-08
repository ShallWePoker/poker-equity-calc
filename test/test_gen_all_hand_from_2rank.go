package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func main() {
	c1 := "A3"
	hcr, err := models.InitHoleCardRange(c1, false)
	if err != nil {
		panic(err)
	}
	holeCards := utils.GenerateAllHandFrom2Ranks(hcr)
	fmt.Printf("holeCards: %+v\n", holeCards)
	fmt.Printf("total: %d\n", len(holeCards))
}
