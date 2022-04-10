package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
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
		fmt.Printf("\nhand1 %s vs hand2 %s\n", hand1.ToString(), hand2Holecard.ToString())
		cardsToRemoveFromDeck := make([]models.Card, 0)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2Holecard...)
		// TODO seems ExamineDistinctCards unnecessary here;
		/*
		err = ExamineDistinctCards(cardsToRemoveFromDeck...)
		if err != nil {
			return hand1Equity, hand2RangeEquity, err
		}
		 */
		defaultSampleSize := 1000
		totalSampleSize = totalSampleSize + defaultSampleSize
		err = Hand2HandCountWinAndTie(hand1, hand2Holecard,
			&hand1Win, &hand2RangesWin, &tie,
			consts.PREFLOP, cardsToRemoveFromDeck, defaultSampleSize,
			true, nil, nil, nil)
		if err != nil {
			return hand1Equity, hand2RangeEquity, err
		}
	}
	hand1WinRate, hand1EquityRate, hand2RangesWinRate, hand2RangesEquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2RangesWin, tie, float64(totalSampleSize))
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2RangeEquity.Equity = hand2RangesEquityRate
	hand2RangeEquity.WinRate = hand2RangesWinRate
	hand2RangeEquity.TieRate = tieRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusRangeFlopEquity(hand1 models.HoleCard, hand2Ranges []models.HoleCardRange, flop models.Flop) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if len(hand2Ranges) == 0 {
		return hand1Equity, hand2RangeEquity, errors.New("hand2 range must not be empty")
	}
	hand1Equity.HoleCard = hand1
	hand2RangeEquity.HoleCardRange = hand2Ranges
	hand1Win := float64(0)
	hand2RangesWin := float64(0)
	tie := float64(0)
	totalSampleSize := 0

	hand2Holecards := GenerateAllHoleCardFromMultiRange(hand2Ranges, append(flop, hand1...)...)

	for _, hand2Holecard := range hand2Holecards {
		fmt.Printf("\nhand1 %s vs hand2 %s\n", hand1.ToString(), hand2Holecard.ToString())
		cardsToRemoveFromDeck := make([]models.Card, 0)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2Holecard...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
		defaultSampleSize := 1000
		totalSampleSize = totalSampleSize + defaultSampleSize
		err = Hand2HandCountWinAndTie(hand1, hand2Holecard,
			&hand1Win, &hand2RangesWin, &tie,
			consts.FLOP, cardsToRemoveFromDeck, defaultSampleSize,
			true, &flop, nil, nil,
		)
	}
	hand1WinRate, hand1EquityRate, hand2RangesWinRate, hand2RangesEquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2RangesWin, tie, float64(totalSampleSize))
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2RangeEquity.Equity = hand2RangesEquityRate
	hand2RangeEquity.WinRate = hand2RangesWinRate
	hand2RangeEquity.TieRate = tieRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusRangeTurnEquity(hand1 models.HoleCard, hand2Ranges []models.HoleCardRange, flop models.Flop, turn models.Card) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if len(hand2Ranges) == 0 {
		return hand1Equity, hand2RangeEquity, errors.New("hand2 range must not be empty")
	}
	hand1Equity.HoleCard = hand1
	hand2RangeEquity.HoleCardRange = hand2Ranges
	hand1Win := float64(0)
	hand2RangesWin := float64(0)
	tie := float64(0)
	totalSampleSize := 0

	hand2Holecards := GenerateAllHoleCardFromMultiRange(hand2Ranges, append(append(flop, turn), hand1...)...)

	for _, hand2Holecard := range hand2Holecards {
		fmt.Printf("\nhand1 %s vs hand2 %s\n", hand1.ToString(), hand2Holecard.ToString())
		cardsToRemoveFromDeck := make([]models.Card, 0)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2Holecard...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, turn)
		defaultSampleSize := 1000
		totalSampleSize += defaultSampleSize
		err = Hand2HandCountWinAndTie(hand1, hand2Holecard,
			&hand1Win, &hand2RangesWin, &tie,
			consts.TURN, cardsToRemoveFromDeck, defaultSampleSize,
			true, &flop, &turn, nil)
		if err != nil {
			return hand1Equity, hand2RangeEquity, err
		}
	}
	hand1WinRate, hand1EquityRate, hand2RangesWinRate, hand2RangesEquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2RangesWin, tie, float64(totalSampleSize))
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2RangeEquity.Equity = hand2RangesEquityRate
	hand2RangeEquity.WinRate = hand2RangesWinRate
	hand2RangeEquity.TieRate = tieRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusRangeRiverEquity(hand1 models.HoleCard, hand2Ranges []models.HoleCardRange, flop models.Flop, turn models.Card, river models.Card) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if len(hand2Ranges) == 0 {
		return hand1Equity, hand2RangeEquity, errors.New("hand2 range must not be empty")
	}
	hand1Equity.HoleCard = hand1
	hand2RangeEquity.HoleCardRange = hand2Ranges
	hand1Win := float64(0)
	hand2RangesWin := float64(0)
	tie := float64(0)
	totalSampleSize := 0

	hand2Holecards := GenerateAllHoleCardFromMultiRange(hand2Ranges, append(append(flop, turn, river), hand1...)...)

	for _, hand2Holecard := range hand2Holecards {
		fmt.Printf("\nhand1 %s vs hand2 %s\n", hand1.ToString(), hand2Holecard.ToString())
		cardsToRemoveFromDeck := make([]models.Card, 0)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2Holecard...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, turn)
		cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, river)
		defaultSampleSize := 1000
		totalSampleSize += defaultSampleSize
		err = Hand2HandCountWinAndTie(hand1, hand2Holecard,
			&hand1Win, &hand2RangesWin, &tie,
			consts.RIVER, cardsToRemoveFromDeck, defaultSampleSize,
			true, &flop, &turn, &river)
		if err != nil {
			return hand1Equity, hand2RangeEquity, err
		}
	}
	hand1WinRate, hand1EquityRate, hand2RangesWinRate, hand2RangesEquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2RangesWin, tie, float64(totalSampleSize))
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2RangeEquity.Equity = hand2RangesEquityRate
	hand2RangeEquity.WinRate = hand2RangesWinRate
	hand2RangeEquity.TieRate = tieRate
	return hand1Equity, hand2RangeEquity, nil
}

