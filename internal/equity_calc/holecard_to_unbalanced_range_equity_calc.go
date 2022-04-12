package equity_calc

import (
	"errors"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"github.com/ShallWePoker/poker-equity-calc/internal/utils"
)

func HoleCardVersusUnbalancedRangePreflopEquity(hand1 models.HoleCard, hand2Ranges []models.UnbalancedHolecardRange) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if err = examineUnbalancedRangesInput(hand2Ranges); err != nil {
		return hand1Equity, hand2RangeEquity, err
	}
	hand1EquityRate := float64(0)
	hand2RangeEquityRate := float64(0)
	for _, hand2RangePart := range hand2Ranges {
		currHand1EquityRate := float64(0)
		currHand2RangeEquityRate := float64(0)
		hand2RangeTotalCombos := 0
		for _, hand2Range := range hand2RangePart.HolecardRanges {
			hand2HolecardCombos := utils.GenerateAllHoleCardFromMultiRange([]models.HoleCardRange{hand2Range}, hand1...)
			hand2HolecardCombosNum := len(hand2HolecardCombos)
			hand1EquityShare, hand2HolecardEquityShare, err := HoleCardVersusBalancedRangePreflopEquity(hand1, []models.HoleCardRange{hand2Range})
			if err != nil {
				return hand1Equity, hand2RangeEquity, err
			}
			currHand1EquityRate += hand1EquityShare.Equity*float64(hand2HolecardCombosNum)
			currHand2RangeEquityRate += hand2HolecardEquityShare.Equity*float64(hand2HolecardCombosNum)
			hand2RangeTotalCombos += hand2HolecardCombosNum
		}
		hand1EquityRate += (currHand1EquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
		hand2RangeEquityRate += (currHand2RangeEquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
	}
	// output equity only; win rate & tie rate are not as important
	hand1Equity.Equity = hand1EquityRate
	hand2RangeEquity.Equity = hand2RangeEquityRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusUnbalancedRangeFlopEquity(hand1 models.HoleCard, hand2Ranges []models.UnbalancedHolecardRange, flop models.Flop) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if err = examineUnbalancedRangesInput(hand2Ranges); err != nil {
		return hand1Equity, hand2RangeEquity, err
	}
	hand1EquityRate := float64(0)
	hand2RangeEquityRate := float64(0)
	for _, hand2RangePart := range hand2Ranges {
		currHand1EquityRate := float64(0)
		currHand2RangeEquityRate := float64(0)
		hand2RangeTotalCombos := 0
		for _, hand2Range := range hand2RangePart.HolecardRanges {
			hand2HolecardCombos := utils.GenerateAllHoleCardFromMultiRange([]models.HoleCardRange{hand2Range}, append(flop, hand1...)...)
			hand2HolecardCombosNum := len(hand2HolecardCombos)
			hand1EquityShare, hand2HolecardEquityShare, err := HoleCardVersusBalancedRangeFlopEquity(hand1, []models.HoleCardRange{hand2Range}, flop)
			if err != nil {
				return hand1Equity, hand2RangeEquity, err
			}
			currHand1EquityRate += hand1EquityShare.Equity*float64(hand2HolecardCombosNum)
			currHand2RangeEquityRate += hand2HolecardEquityShare.Equity*float64(hand2HolecardCombosNum)
			hand2RangeTotalCombos += hand2HolecardCombosNum
		}
		hand1EquityRate += (currHand1EquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
		hand2RangeEquityRate += (currHand2RangeEquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
	}
	// output equity only; win rate & tie rate are not as important
	hand1Equity.Equity = hand1EquityRate
	hand2RangeEquity.Equity = hand2RangeEquityRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusUnbalancedRangeTurnEquity(hand1 models.HoleCard, hand2Ranges []models.UnbalancedHolecardRange, flop models.Flop, turn models.Card) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if err = examineUnbalancedRangesInput(hand2Ranges); err != nil {
		return hand1Equity, hand2RangeEquity, err
	}
	hand1EquityRate := float64(0)
	hand2RangeEquityRate := float64(0)
	for _, hand2RangePart := range hand2Ranges {
		currHand1EquityRate := float64(0)
		currHand2RangeEquityRate := float64(0)
		hand2RangeTotalCombos := 0
		for _, hand2Range := range hand2RangePart.HolecardRanges {
			hand2HolecardCombos := utils.GenerateAllHoleCardFromMultiRange([]models.HoleCardRange{hand2Range}, append(append(flop, turn), hand1...)...)
			hand2HolecardCombosNum := len(hand2HolecardCombos)
			hand1EquityShare, hand2HolecardEquityShare, err := HoleCardVersusBalancedRangeTurnEquity(hand1, []models.HoleCardRange{hand2Range}, flop, turn)
			if err != nil {
				return hand1Equity, hand2RangeEquity, err
			}
			currHand1EquityRate += hand1EquityShare.Equity*float64(hand2HolecardCombosNum)
			currHand2RangeEquityRate += hand2HolecardEquityShare.Equity*float64(hand2HolecardCombosNum)
			hand2RangeTotalCombos += hand2HolecardCombosNum
		}
		hand1EquityRate += (currHand1EquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
		hand2RangeEquityRate += (currHand2RangeEquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
	}
	// output equity only; win rate & tie rate are not as important
	hand1Equity.Equity = hand1EquityRate
	hand2RangeEquity.Equity = hand2RangeEquityRate
	return hand1Equity, hand2RangeEquity, nil
}

func HoleCardVersusUnbalancedRangeRiverEquity(hand1 models.HoleCard, hand2Ranges []models.UnbalancedHolecardRange, flop models.Flop, turn models.Card, river models.Card) (hand1Equity models.HoleCardEquity, hand2RangeEquity models.HoleCardRangeEquity, err error) {
	if err = examineUnbalancedRangesInput(hand2Ranges); err != nil {
		return hand1Equity, hand2RangeEquity, err
	}
	hand1EquityRate := float64(0)
	hand2RangeEquityRate := float64(0)
	for _, hand2RangePart := range hand2Ranges {
		currHand1EquityRate := float64(0)
		currHand2RangeEquityRate := float64(0)
		hand2RangeTotalCombos := 0
		for _, hand2Range := range hand2RangePart.HolecardRanges {
			hand2HolecardCombos := utils.GenerateAllHoleCardFromMultiRange([]models.HoleCardRange{hand2Range}, append(append(flop, turn, river), hand1...)...)
			hand2HolecardCombosNum := len(hand2HolecardCombos)
			hand1EquityShare, hand2HolecardEquityShare, err := HoleCardVersusBalancedRangeRiverEquity(hand1, []models.HoleCardRange{hand2Range}, flop, turn, river)
			if err != nil {
				return hand1Equity, hand2RangeEquity, err
			}
			currHand1EquityRate += hand1EquityShare.Equity*float64(hand2HolecardCombosNum)
			currHand2RangeEquityRate += hand2HolecardEquityShare.Equity*float64(hand2HolecardCombosNum)
			hand2RangeTotalCombos += hand2HolecardCombosNum
		}
		hand1EquityRate += (currHand1EquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
		hand2RangeEquityRate += (currHand2RangeEquityRate/float64(hand2RangeTotalCombos))*(hand2RangePart.Percentage)
	}
	// output equity only; win rate & tie rate are not as important
	hand1Equity.Equity = hand1EquityRate
	hand2RangeEquity.Equity = hand2RangeEquityRate
	return hand1Equity, hand2RangeEquity, nil
}

func examineUnbalancedRangesInput(unbalancedRanges []models.UnbalancedHolecardRange) error {
	totalPercentage := float64(0)
	for _, range1 := range unbalancedRanges {
		totalPercentage += range1.Percentage
	}
	if totalPercentage != 1.0 {
		return errors.New("unbalanced ranges percentages do not sum up to 1.0")
	}
	// TODO needed more check if unbalancedRanges overlap each other; or we have to depend on function caller's input
	/*
		unbalancedRanges:
				[AA, KK]  90%  +  [KK, QQ] 10%
				KK overlap
	 */
	return nil
}