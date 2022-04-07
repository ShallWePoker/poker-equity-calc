package main

import (
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func main() {
	c1 := "2"
	c2 := "2"
	holeCards := utils.GenerateAllHandFrom2Ranks(c1, c2)
	fmt.Printf("holeCards: %+v\n", holeCards)
	fmt.Printf("total: %d\n", len(holeCards))
}
