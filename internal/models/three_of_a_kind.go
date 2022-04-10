package models

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
)

type ThreeOfAKind struct {
	Hand          Hand
	ThreeCardRank int
	HighCardRank  int
	LowCardRank   int
}

func (t ThreeOfAKind) ToString() string {
	return t.Hand.ToString()
}

func (t ThreeOfAKind) BoardCards() Hand {
	return t.Hand
}

func (t ThreeOfAKind) Category() string {
	return consts.ThreeOfAKind
}

func (t ThreeOfAKind) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse ||
		madeHand.Category() == consts.Flush ||
		madeHand.Category() == consts.Straight {
		return false
	} else if madeHand.Category() != consts.ThreeOfAKind {
		return true
	} else {
		anotherHand := madeHand.(ThreeOfAKind)
		if t.ThreeCardRank != anotherHand.ThreeCardRank {
			return t.ThreeCardRank - anotherHand.ThreeCardRank > 0
		} else {
			tValue := t.HighCardRank*10+t.LowCardRank
			aValue := anotherHand.HighCardRank*10+anotherHand.LowCardRank
			return tValue- aValue > 0
		}
	}
}
