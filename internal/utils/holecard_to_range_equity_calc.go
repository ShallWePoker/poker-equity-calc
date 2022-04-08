package utils

import (
	"errors"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func HoleCardVersusRangePreflopEquity(hand1 models.HoleCard, hand2Ranges []models.HoleCardRange) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if len(hand2Ranges) == 0 {
		return hand1Equity, hand2RangeEquity, errors.New("hand2 range must not be empty")
	}
	hand1Equity.HoleCard = hand1
	hand2RangeEquity.HoleCardRange = hand2Ranges

	hand1Win := float64(0)
	hand2RangesWin := float64(0)
	tie := float64(0)
	totalSampleSize := 0

	hand2Holecards := GenerateAllHoleCardFromMultiRange(hand2Ranges, hand1...)

	for _, hand2Holecard := range hand2Holecards {
		cardsToRemoveFromDeck := make([]models.Card, 0)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2Holecard...)
		// TODO seems ExamineDistinctCards unnecessary here;
		err = ExamineDistinctCards(cardsToRemoveFromDeck...)
		if err != nil {
			return hand1Equity, hand2RangeEquity, err
		}
		defaultSampleSize := 500
		totalSampleSize = totalSampleSize + defaultSampleSize
		for i := 0; i < defaultSampleSize; i++ {
			generated5Cards, err := GenerateRandomNCards(5, cardsToRemoveFromDeck)
			if err != nil {
				return hand1Equity, hand2RangeEquity, err
			}
			hand1BiggestMadeHand := Seven2five(append(generated5Cards, hand1...))
			hand2BiggestMadeHand := Seven2five(append(generated5Cards, hand2Holecard...))
			if hand1BiggestMadeHand.IsGreaterThan(hand2BiggestMadeHand) {
				hand1Win += 1
			} else if hand2BiggestMadeHand.IsGreaterThan(hand1BiggestMadeHand) {
				hand2RangesWin += 1
			} else {
				tie += 1
			}
		}
	}
	hand1WinRate := hand1Win / float64(totalSampleSize)
	hand1EquityRate := (hand1Win+tie/float64(2)) / float64(totalSampleSize)
	hand2RangesWinRate := hand2RangesWin / float64(totalSampleSize)
	hand2RangesEquityRate := (hand2RangesWin+tie/float64(2)) / float64(totalSampleSize)
	tieRate := tie / float64(totalSampleSize)
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2RangeEquity.Equity = hand2RangesEquityRate
	hand2RangeEquity.WinRate = hand2RangesWinRate
	hand2RangeEquity.TieRate = tieRate
	return hand1Equity, hand2RangeEquity, nil
}
