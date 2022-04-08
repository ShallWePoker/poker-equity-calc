package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func GenerateAllHandFrom2Ranks(hcr models.HoleCardRange) []models.HoleCards {
	if hcr.TitlePair[1] == hcr.TitlePair[0] {
		return generatePairHand(hcr)
	} else {
		return generateUnPairHand(hcr)
	}
}

func generatePairHand(c models.HoleCardRange) []models.HoleCards {
	var card1, card2 models.Card
	var holeCardList []models.HoleCards
	var err error
	for i := 0; i < 4; i++ {
		card1, err = models.InitCardFromString(string(c.TitlePair[0]) + consts.Suits[i])
		if err != nil {
			panic(err)
		}
		for j := i + 1; j < 4; j++ {
			card2, err = models.InitCardFromString(string(c.TitlePair[0]) + consts.Suits[j])
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

func generateUnPairHand(hcr models.HoleCardRange) []models.HoleCards {
	var card1, card2 models.Card
	var holeCardList []models.HoleCards
	var err error
	if hcr.Suited == true {
		for i := 0; i < 4; i++ {
			card1, err = models.InitCardFromString(string(hcr.TitlePair[0]) + consts.Suits[i])
			card2, err = models.InitCardFromString(string(hcr.TitlePair[1]) + consts.Suits[i])
			holeCard := models.HoleCards{}
			holeCard, err = models.InitHoleCards([]models.Card{card1, card2})
			if err != nil {
				panic(err)
			}
			holeCardList = append(holeCardList, holeCard)
		}
	} else if hcr.Suited == false {
		for i := 0; i < 4; i++ {
			card1, err = models.InitCardFromString(string(hcr.TitlePair[0]) + consts.Suits[i])
			if err != nil {
				panic(err)
			}
			for j := 0; j < 4; j++ {
				if j == i {
					continue
				}
				card2, err = models.InitCardFromString(string(hcr.TitlePair[1]) + consts.Suits[j])
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
	}

	return holeCardList
}
