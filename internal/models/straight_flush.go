package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type StraightFlush struct {
	Hand Hand
}

func (sf StraightFlush) ToString() string {
	return sf.Hand.ToString()
}

func (sf StraightFlush) BoardCards() Hand {
	return sf.Hand
}

func (sf StraightFlush) Category() string {
	return consts.StraightFlush
}

func (sf StraightFlush) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() != consts.StraightFlush {
		return true
	} else {
		anotherHand := madeHand.(StraightFlush)
		if sf.Hand[4].Rank == 14 || anotherHand.Hand[4].Rank == 14 {
			return sf.Hand[0].Rank - anotherHand.Hand[0].Rank > 0
		}
		return sf.Hand[4].Rank - anotherHand.Hand[4].Rank > 0
	}
}