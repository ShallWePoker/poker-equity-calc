package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type HighCard struct {
	Hand Hand
}

func (h HighCard) ToString() string {
	return h.Hand.ToString()
}

func (h HighCard) BoardCards() Hand {
	return h.Hand
}

func (h HighCard) Category() string {
	return consts.HighCard
}

func (h HighCard) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse ||
		madeHand.Category() == consts.Flush ||
		madeHand.Category() == consts.ThreeOfAKind ||
		madeHand.Category() == consts.TwoPairs ||
		madeHand.Category() == consts.OnePair {
		return false
	} else {
		anotherHand := madeHand.(HighCard)
		hValue := h.Hand[4].Rank*10000+h.Hand[3].Rank*1000+h.Hand[2].Rank*100+h.Hand[1].Rank*10+h.Hand[0].Rank
		aValue := anotherHand.Hand[4].Rank*10000+anotherHand.Hand[3].Rank*1000+anotherHand.Hand[2].Rank*100+anotherHand.Hand[1].Rank*10+anotherHand.Hand[0].Rank
		return hValue - aValue > 0
	}
}
