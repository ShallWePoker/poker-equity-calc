package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type TwoPairs struct {
	Hand            Hand
	HighTwoCardRank int
	LowTwoCardRank  int
	SingleCardRank  int
}

func (t TwoPairs) ToString() string {
	return t.Hand.ToString()
}

func (t TwoPairs) BoardCards() Hand {
	return t.Hand
}

func (t TwoPairs) Category() string {
	return consts.TwoPairs
}

func (t TwoPairs) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse ||
		madeHand.Category() == consts.Flush ||
		madeHand.Category() == consts.ThreeOfAKind {
		return false
	} else if madeHand.Category() != consts.TwoPairs {
		return true
	} else {
		anotherHand := madeHand.(TwoPairs)
		tValue := t.HighTwoCardRank*100+t.LowTwoCardRank*10+t.SingleCardRank
		aValue := anotherHand.HighTwoCardRank*100+anotherHand.LowTwoCardRank*10+anotherHand.SingleCardRank
		return tValue - aValue > 0
	}
}

