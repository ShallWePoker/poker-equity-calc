package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func Seven2five(rawCards []models.Card) models.MadeHand {
	var MadeHandList []models.MadeHand
	removed2CardsIndicesList := make([][]int, 0)
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 7; j++ {
			pair := []int{i, j}
			removed2CardsIndicesList = append(removed2CardsIndicesList, pair)
		}
	}
	for _, removed2CardsIndices := range removed2CardsIndicesList {
		cards := make([]models.Card, 0)
		for i, card := range rawCards {
			if !elementIn(i, removed2CardsIndices) {
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

func Seven2fiveV2(rawCards []models.Card) models.MadeHand {
	return FindBiggestMadeHand(GetCombineMatch(rawCards, 5))
}

func combineLoop(arr []models.Card, r []models.Card, i int, n int, output chan<- []models.Card) {
	if n <= 0 {
		return
	}
	rlen := len(r) - n
	alen := len(arr)
	for j := i; j < alen; j++ {
		r[rlen] = arr[j]
		if n == 1 {
			or := make([]models.Card, len(r))
			copy(or, r)
			output <- or
		} else {
			combineLoop(arr, r, j+1, n-1, output)
		}
	}
}

func GetCombineMatch(arr []models.Card, n int) []models.MadeHand {
	var MadeHandList []models.MadeHand
	output := make(chan []models.Card)
	r := make([]models.Card, n)
	go func() {
		combineLoop(arr, r, 0, n, output)
		close(output)
	}()
	for arr := range output {
		hand, _ := models.InitHand(arr)
		madeHand := hand.Categorize()
		MadeHandList = append(MadeHandList, madeHand)
	}

	return MadeHandList
}
