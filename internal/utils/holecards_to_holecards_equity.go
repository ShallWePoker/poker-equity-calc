package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func HandVersusHandPreflopEquity(hand1 models.HoleCards, hand2 models.HoleCards) (hand1Equity models.HoleCardsEquity, hand2Equity models.HoleCardsEquity, err error) {
	hand1Equity.HoleCards = hand1
	hand2Equity.HoleCards = hand2
	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2...)
	err = ExamineDistinctCards(cardsToRemoveFromDeck...)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	var hand1Win float64 = 0
	var hand2Win float64 = 0
	tie := 0
	defaultSampleSize := 1000
	defaultInspectSamplePerHandSize := 42
	inspectSample := false
	for i := 0; i < defaultSampleSize; i++ {
		generated5Cards, err := GenerateRandomNCards(5, cardsToRemoveFromDeck)
		if err != nil {
			return hand1Equity, hand2Equity, err
		}
		if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
			fmt.Printf("board is: %s\n", models.Hand(generated5Cards).ToString())
		}
		hand1BiggestMadeHand := Seven2five(append(generated5Cards, hand1...))
		hand2BiggestMadeHand := Seven2five(append(generated5Cards, hand2...))
		if hand1BiggestMadeHand.IsGreaterThan(hand2BiggestMadeHand) {
			hand1Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s wins hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else if hand2BiggestMadeHand.IsGreaterThan(hand1BiggestMadeHand) {
			hand2Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s loses hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else {
			tie += 1
			hand1Win += 0.5
			hand2Win += 0.5
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s ties hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		}
	}
	hand1WinRate := hand1Win / float64(defaultSampleSize)
	hand2WinRate := hand2Win / float64(defaultSampleSize)
	tieRate := float64(tie) / float64(defaultSampleSize)
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
	hand2Equity.TieRate = tieRate
	return hand1Equity, hand2Equity, nil
}

func HandVersusHandFlopEquity(hand1 models.HoleCards, hand2 models.HoleCards, flop models.Flop) (hand1Equity models.HoleCardsEquity, hand2Equity models.HoleCardsEquity, err error) {
	hand1Equity.HoleCards = hand1
	hand2Equity.HoleCards = hand2
	cardsToRemoveFromDeck := make([]models.Card, 0)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand1...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, hand2...)
	cardsToRemoveFromDeck = append(cardsToRemoveFromDeck, flop...)
	err = ExamineDistinctCards(cardsToRemoveFromDeck...)
	if err != nil {
		return hand1Equity, hand2Equity, err
	}
	var hand1Win float64 = 0
	var hand2Win float64 = 0
	var tie float64 = 0
	defaultSampleSize := 1000
	defaultInspectSamplePerHandSize := 42
	inspectSample := false
	for i := 0; i < defaultSampleSize; i++ {
		generated2Cards, err := GenerateRandomNCards(2, cardsToRemoveFromDeck)
		if err != nil {
			return hand1Equity, hand2Equity, err
		}
		board := make([]models.Card, 0)
		board = append(board, flop...)
		board = append(board, generated2Cards...)
		if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
			fmt.Printf("board is: %s\n", models.Hand(board).ToString())
		}
		hand1BiggestMadeHand := Seven2five(append(board, hand1...))
		hand2BiggestMadeHand := Seven2five(append(board, hand2...))
		if hand1BiggestMadeHand.IsGreaterThan(hand2BiggestMadeHand) {
			hand1Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s wins hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else if hand2BiggestMadeHand.IsGreaterThan(hand1BiggestMadeHand) {
			hand2Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s loses hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else {
			tie += 1
			hand1Win += 0.5
			hand2Win += 0.5
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s ties hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		}
	}
	hand1WinRate := hand1Win / float64(defaultSampleSize)
	hand2WinRate := hand2Win / float64(defaultSampleSize)
	tieRate := tie / float64(defaultSampleSize)
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
	hand2Equity.TieRate = tieRate
	return hand1Equity, hand2Equity, nil
}

func HandVersusHandTurnEquity(hand1 models.HoleCards, hand2 models.HoleCards, flop models.Flop, turn models.Card) (hand1Equity models.HoleCardsEquity, hand2Equity models.HoleCardsEquity, err error) {
	hand1Equity.HoleCards = hand1
	hand2Equity.HoleCards = hand2
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
	defaultInspectSamplePerHandSize := 42
	inspectSample := false
	for i := 0; i < defaultSampleSize; i++ {
		generated1Card, err := GenerateRandomNCards(1, cardsToRemoveFromDeck)
		if err != nil {
			return hand1Equity, hand2Equity, err
		}
		board := make([]models.Card, 0)
		board = append(board, flop...)
		board = append(board, turn)
		board = append(board, generated1Card...)
		if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
			fmt.Printf("board is: %s\n", models.Hand(board).ToString())
		}
		hand1BiggestMadeHand := Seven2five(append(board, hand1...))
		hand2BiggestMadeHand := Seven2five(append(board, hand2...))
		if hand1BiggestMadeHand.IsGreaterThan(hand2BiggestMadeHand) {
			hand1Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s wins hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else if hand2BiggestMadeHand.IsGreaterThan(hand1BiggestMadeHand) {
			hand2Win += 1
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s loses hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		} else {
			tie += 1
			hand1Win += 0.5
			hand2Win += 0.5
			if (inspectSample) && (defaultInspectSamplePerHandSize != 0) && (i%defaultInspectSamplePerHandSize == 0) {
				fmt.Printf("hand1's %s ties hand2's %s\n\n", hand1BiggestMadeHand.ToString(), hand2BiggestMadeHand.ToString())
			}
		}
	}
	hand1WinRate := hand1Win / float64(defaultSampleSize)
	hand2WinRate := hand2Win / float64(defaultSampleSize)
	tieRate := tie / float64(defaultSampleSize)
	hand1Equity.WinRate = hand1WinRate
	hand1Equity.TieRate = tieRate
	hand2Equity.WinRate = hand2WinRate
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

