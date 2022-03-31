package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func Seven2five(rawCards []models.Card) models.MadeHand {
	var MadeHandList []models.MadeHand
	removed2CardsIndicesList := make([][]int, 0)
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 7; j++ {
			pair := []int{i,j}
			removed2CardsIndicesList = append(removed2CardsIndicesList, pair)
		}
	}
	for _, removed2CardsIndices := range removed2CardsIndicesList {
		cards := make([]models.Card, 0)
		for i, card := range rawCards {
			if !elementIn(i,removed2CardsIndices) {
				cards = append(cards, card)
			}
		}
		hand, _ := models.InitHand(cards)
		madeHand := hand.Categorize()
		MadeHandList = append(MadeHandList, madeHand)
	}
	return FindBiggestMadeHand(MadeHandList)
}

func elementIn(elem int, bound []int) bool {
	for _, i := range bound {
		if elem == i {
			return true
		}
	}
	return false
}

func FindBiggestMadeHand(MadeHandList []models.MadeHand) models.MadeHand {
	max := MadeHandList[0]
	for i := 1; i < len(MadeHandList); i++ {
		if MadeHandList[i].IsGreaterThan(max) {
			max = MadeHandList[i]
		}
	}
	return max
}
