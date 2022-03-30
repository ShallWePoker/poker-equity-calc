package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type Straight struct {
	Hand Hand
}

func (s Straight) ToString() string {
	return s.Hand.ToString()
}

func (s Straight) BoardCards() Hand {
	return s.Hand
}

func (s Straight) Category() string {
	return consts.Straight
}

func (s Straight) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse ||
		madeHand.Category() == consts.Flush {
		return false
	} else if madeHand.Category() != consts.Straight {
		return true
	} else {
		anotherHand := madeHand.(Straight)
		if s.Hand[4].Rank == 14 || anotherHand.Hand[4].Rank == 14 {
			return s.Hand[0].Rank - anotherHand.Hand[0].Rank > 0
		}
		return s.Hand[4].Rank - anotherHand.Hand[4].Rank > 0
	}

}
