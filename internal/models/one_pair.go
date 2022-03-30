package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type OnePair struct {
	Hand         Hand
	PairCardRank int
	LowCardRank  int
	MidCardRank  int
	HighCardRank int
}

func (o OnePair) ToString() string {
	return o.Hand.ToString()
}

func (o OnePair) BoardCards() Hand {
	return o.Hand
}

func (o OnePair) Category() string {
	return consts.OnePair
}

func (o OnePair) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse ||
		madeHand.Category() == consts.Flush ||
		madeHand.Category() == consts.ThreeOfAKind ||
		madeHand.Category() == consts.TwoPairs {
		return false
	} else if madeHand.Category() != consts.OnePair {
		return true
	} else {
		anotherHand := madeHand.(OnePair)
		oValue := o.PairCardRank*1000+o.HighCardRank*100+o.MidCardRank*10+o.LowCardRank
		aValue := anotherHand.PairCardRank*1000+anotherHand.HighCardRank*100+anotherHand.MidCardRank*10+anotherHand.LowCardRank
		return oValue - aValue > 0
	}
}