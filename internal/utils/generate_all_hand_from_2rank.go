package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func GenerateAllHandFrom2Ranks(c1, c2 string) []models.HoleCards {
	if c1 == c2 {
		return generateSuitedHand(c1)
	} else {
		return generateOffSuitedHand(c1, c2)
	}
}

func generateSuitedHand(c string) []models.HoleCards {
	var card1, card2 models.Card
	var holeCardList []models.HoleCards
	var err error
	for i := 0; i < 4; i++ {
		card1, err = models.InitCardFromString(c + consts.Suits[i])
		if err != nil {
			panic(err)
		}
		for j := i + 1; j < 4; j++ {
			card2, err = models.InitCardFromString(c + consts.Suits[j])
			if err != nil {
				panic(err)
			}
			holeCard := models.HoleCards{}
			holeCard, err = models.InitHoleCards([]models.Card{card1, card2})
			if err != nil {
				panic(err)
			}
			holeCardList = append(holeCardList, holeCard)
		}
	}
	return holeCardList
}

func generateOffSuitedHand(c1, c2 string) []models.HoleCards {
	var card1, card2 models.Card
	var holeCardList []models.HoleCards
	var err error
	for i := 0; i < 4; i++ {
		card1, err = models.InitCardFromString(c1 + consts.Suits[i])
		if err != nil {
			panic(err)
		}
		for j := 0; j < 4; j++ {
			card2, err = models.InitCardFromString(c2 + consts.Suits[j])
			if err != nil {
				panic(err)
			}
			holeCard := models.HoleCards{}
			holeCard, err = models.InitHoleCards([]models.Card{card1, card2})
			if err != nil {
				panic(err)
			}
			holeCardList = append(holeCardList, holeCard)
		}
	}
	return holeCardList
}
