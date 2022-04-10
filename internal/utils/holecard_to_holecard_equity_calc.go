package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func HandVersusHandPreflopEquity(hand1 models.HoleCard, hand2 models.HoleCard) (hand1Equity models.HoleCardEquity, hand2Equity models.HoleCardEquity, err error) {
	hand1Equity.HoleCard = hand1
	hand2Equity.HoleCard = hand2
	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2...)
	err = ExamineDistinctCards(cardsToRemoveFromDeck...)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1Win := float64(0)
	hand2Win := float64(0)
	tie := float64(0)
	defaultSampleSize := 1000
	err = Hand2HandCountWinAndTie(hand1, hand2, &hand1Win, &hand2Win, &tie, consts.PREFLOP, cardsToRemoveFromDeck, defaultSampleSize, true, nil, nil, nil)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1WinRate, hand1EquityRate, hand2WinRate, hand2EquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2Win, tie, float64(defaultSampleSize))
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
	hand2Equity.Equity = hand2EquityRate
	hand2Equity.TieRate = tieRate
	return hand1Equity, hand2Equity, nil
}

func HandVersusHandFlopEquity(hand1 models.HoleCard, hand2 models.HoleCard, flop models.Flop) (hand1Equity models.HoleCardEquity, hand2Equity models.HoleCardEquity, err error) {
	hand1Equity.HoleCard = hand1
	hand2Equity.HoleCard = hand2
	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
	err = ExamineDistinctCards(cardsToRemoveFromDeck...)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1Win := float64(0)
	hand2Win := float64(0)
	tie := float64(0)
	defaultSampleSize := 1000
	err = Hand2HandCountWinAndTie(hand1, hand2, &hand1Win, &hand2Win, &tie, consts.FLOP, cardsToRemoveFromDeck, defaultSampleSize, true, &flop, nil, nil)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1WinRate, hand1EquityRate, hand2WinRate, hand2EquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2Win, tie, float64(defaultSampleSize))
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
	hand2Equity.Equity = hand2EquityRate
	hand2Equity.TieRate = tieRate
	return hand1Equity, hand2Equity, nil
}

func HandVersusHandTurnEquity(hand1 models.HoleCard, hand2 models.HoleCard, flop models.Flop, turn models.Card) (hand1Equity models.HoleCardEquity, hand2Equity models.HoleCardEquity, err error) {
	hand1Equity.HoleCard = hand1
	hand2Equity.HoleCard = hand2
	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, turn)
	err = ExamineDistinctCards(cardsToRemoveFromDeck...)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1Win := float64(0)
	hand2Win := float64(0)
	tie := float64(0)
	defaultSampleSize := 1000
	err = Hand2HandCountWinAndTie(hand1, hand2, &hand1Win, &hand2Win, &tie, consts.TURN, cardsToRemoveFromDeck, defaultSampleSize, true, &flop, &turn, nil)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	hand1WinRate, hand1EquityRate, hand2WinRate, hand2EquityRate, tieRate := Hand2HandRateCalc(hand1Win, hand2Win, tie, float64(defaultSampleSize))
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.Equity = hand1EquityRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
	hand2Equity.Equity = hand2EquityRate
	hand2Equity.TieRate = tieRate
	return hand1Equity, hand2Equity, nil
}

func ExamineDistinctCards(cards ...models.Card) error {
	distinctCards := make(map[models.Card]struct{})
	for _, card := range cards {
		_, exists := distinctCards[card]
		if !exists {
			distinctCards[card] = struct{}{}
		} else {
			return errors.New(fmt.Sprintf("duplicate card: %s", card.ToString()))
		}
	}
	return nil
}

