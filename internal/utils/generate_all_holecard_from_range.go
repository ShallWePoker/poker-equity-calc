package utils

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"strings"
)

func GenerateAllHoleCardFromSingleRange(holeCardRange models.HoleCardRange) []models.HoleCard {
	if holeCardRange.TitlePair[1] == holeCardRange.TitlePair[0] {
		return generatePairHand(holeCardRange)
	} else {
		return generateUnPairHand(holeCardRange)
	}
}

func generatePairHand(holeCardRange models.HoleCardRange) []models.HoleCard {
	var card1, card2 models.Card
	var holeCardList []models.HoleCard
	var err error
	for i := 0; i < 4; i++ {
		card1, err = models.InitCardFromString(string(holeCardRange.TitlePair[0]) + consts.Suits[i])
		if err != nil {
			panic(err)
		}
		for j := i + 1; j < 4; j++ {
			card2, err = models.InitCardFromString(string(holeCardRange.TitlePair[0]) + consts.Suits[j])
			if err != nil {
				panic(err)
			}
			holeCard := models.HoleCard{}
			holeCard, err = models.InitHoleCard([]models.Card{card1, card2})
			if err != nil {
				panic(err)
			}
			holeCardList = append(holeCardList, holeCard)
		}
	}
	return holeCardList
}

func generateUnPairHand(holeCardRange models.HoleCardRange) []models.HoleCard {
	var card1, card2 models.Card
	var holeCardList []models.HoleCard
	var err error
	if holeCardRange.Suited == true {
		for i := 0; i < 4; i++ {
			card1, err = models.InitCardFromString(string(holeCardRange.TitlePair[0]) + consts.Suits[i])
			card2, err = models.InitCardFromString(string(holeCardRange.TitlePair[1]) + consts.Suits[i])
			holeCard := models.HoleCard{}
			holeCard, err = models.InitHoleCard([]models.Card{card1, card2})
			if err != nil {
				panic(err)
			}
			holeCardList = append(holeCardList, holeCard)
		}
	} else if holeCardRange.Suited == false {
		for i := 0; i < 4; i++ {
			card1, err = models.InitCardFromString(string(holeCardRange.TitlePair[0]) + consts.Suits[i])
			if err != nil {
				panic(err)
			}
			for j := 0; j < 4; j++ {
				if j == i {
					continue
				}
				card2, err = models.InitCardFromString(string(holeCardRange.TitlePair[1]) + consts.Suits[j])
				if err != nil {
					panic(err)
				}
				holeCard := models.HoleCard{}
				holeCard, err = models.InitHoleCard([]models.Card{card1, card2})
				if err != nil {
					panic(err)
				}
				holeCardList = append(holeCardList, holeCard)
			}
		}
	}

	return holeCardList
}

func GenerateAllHoleCardFromMultiRange(holeCardRanges []models.HoleCardRange, excludedCards ...models.Card) (resp []models.HoleCard) {
	holeCardsMap := make(map[string]bool)
	for _, holeCardRange := range holeCardRanges {
		allHoleCard := GenerateAllHoleCardFromSingleRange(holeCardRange)
		for _, holeCard := range allHoleCard {
			if _, exists := holeCardsMap[holeCard.Key()]; !exists {
				holeCardsMap[holeCard.Key()] = true
			}
		}
	}
	for value, _ := range holeCardsMap {
		valueSlice := strings.Split(value, " ")
		card1Str := valueSlice[0]
		card2Str := valueSlice[1]
		card1, err := models.InitCardFromString(card1Str)
		if err != nil {
			panic(err)
		}
		card2, err := models.InitCardFromString(card2Str)
		if err != nil {
			panic(err)
		}
		hc, err := models.InitHoleCard([]models.Card{card1, card2})
		if err != nil {
			panic(err)
		}
		resp = append(resp, hc)
	}
	return PickOutBlockedHoleCard(resp, excludedCards...)
}

func PickOutBlockedHoleCard(holeCards []models.HoleCard, excludedCards ...models.Card) (resp []models.HoleCard) {
	for _, holecard := range holeCards {
		if !holecard.BlockedBy(excludedCards...) {
			resp = append(resp, holecard)
		}
	}
	return resp
}

