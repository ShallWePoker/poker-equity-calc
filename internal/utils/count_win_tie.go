package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func Hand2HandCountWinAndTie(
	hand1, hand2 models.HoleCard,
	hand1WinCount, hand2WinCount, tieCount *float64,
	gameStage string,
	cardsToRemoveFromDeck []models.Card,
	sampleSize int, inspectSample bool,
	flop *models.Flop, turn *models.Card, river *models.Card) error {
	if inspectSample {
		fmt.Printf("\nhand1 %s vs hand2 %s\n", hand1.ToString(), hand2.ToString())
	}
	defaultInspectSamplePerHandSize := 42
	if sampleSize == 0 {
		sampleSize = 1000
	}
	unknownCardNum, exists := consts.STAGE_UNKNOWN_CARD_NUM[gameStage]
	if !exists {
		return errors.New("invalid game stage: " + gameStage)
	}
	for i := 0; i < sampleSize; i++ {
		var generatedCards []models.Card
		var err error
		if gameStage != consts.RIVER {
			generatedCards, err = GenerateRandomNCards(unknownCardNum, cardsToRemoveFromDeck)
			if err != nil {
				return err
			}
		}
		var board []models.Card
		if gameStage == consts.PREFLOP {
			board = append(board, generatedCards...)
		} else if gameStage == consts.FLOP {
			board = append(board, *flop...)
			board = append(board, generatedCards...)
		} else if gameStage == consts.TURN {
			board = append(board, *flop...)
			board = append(board, *turn)
			board = append(board, generatedCards...)
		} else {
			board = append(board, *flop...)
			board = append(board, *turn)
			board = append(board, *river)
		}
		if inspectSample && (i%defaultInspectSamplePerHandSize == 0) {
			fmt.Printf("\nboard is: %s\n", models.Hand(board).ToString())
		}
		hand1MadeHand := Seven2five(append(board, hand1...))
		hand2MadeHand := Seven2five(append(board, hand2...))
		if inspectSample && (i%defaultInspectSamplePerHandSize == 0) {
			fmt.Printf("hand1MadeHand %s is %s\n", hand1MadeHand.ToString(), hand1MadeHand.Category())
			fmt.Printf("hand2MadeHand %s is %s\n", hand2MadeHand.ToString(), hand2MadeHand.Category())
		}
		if hand1MadeHand.IsGreaterThan(hand2MadeHand) {
			*hand1WinCount += 1
			if inspectSample && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s wins hand2's %s\n", hand1MadeHand.ToString(), hand2MadeHand.ToString())
			}
		} else if hand2MadeHand.IsGreaterThan(hand1MadeHand) {
			*hand2WinCount += 1
			if inspectSample && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s loses hand2's %s\n", hand1MadeHand.ToString(), hand2MadeHand.ToString())
			}
		} else {
			*tieCount += 1
			if inspectSample && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s ties hand2's %s\n", hand1MadeHand.ToString(), hand2MadeHand.ToString())
			}
		}
	}
	return nil
}

func Hand2HandRateCalc(hand1WinCount, hand2WinCount, tieCount float64, totalSampleSize float64) (hand1WinRate, hand1EquityRate, hand2WinRate, hand2EquityRate, tieRate float64) {
	hand1WinRate = hand1WinCount / totalSampleSize
	hand1EquityRate = (hand1WinCount + tieCount/float64(2)) / totalSampleSize
	hand2WinRate = hand2WinCount / totalSampleSize
	hand2EquityRate = (hand2WinCount + tieCount/float64(2)) / totalSampleSize
	tieRate = tieCount / totalSampleSize
	return hand1WinRate, hand1EquityRate, hand2WinRate, hand2EquityRate, tieRate
}
