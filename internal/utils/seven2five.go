package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func Seven2five(rawCards []models.Card) models.MadeHand {
	var MadeHandList []models.MadeHand
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 6; j++ {
			var cards []models.Card
			cards = append(rawCards[:i], rawCards[i+1:j]...)
			cards = append(cards, rawCards[j+1:]...)
			hand, err := models.InitHand(cards)
			if err != nil {
				panic(err)
			}
			// fmt.Println(hand.ToString())
			madeHand := hand.Categorize()
			MadeHandList = append(MadeHandList, madeHand)
		}
	}
	return findMax(MadeHandList)
}

func findMax(MadeHandList []models.MadeHand) models.MadeHand {
	max := MadeHandList[0]
	for i := 1; i < len(MadeHandList); i++ {
		if MadeHandList[i].IsGreaterThan(max) {
			max = MadeHandList[i]
		}
	}
	return max
}
