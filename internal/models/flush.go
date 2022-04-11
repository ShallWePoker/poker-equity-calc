package models

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
)

type Flush struct {
	Hand Hand
}

var _ MadeHand = (*Flush)(nil)

func (f Flush) ToString() string {
	return f.Hand.ToString()
}

func (f Flush) BoardCards() Hand {
	return f.Hand
}

func (f Flush) Category() string {
	return consts.Flush
}

func (f Flush) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush ||
		madeHand.Category() == consts.FourOfAKind ||
		madeHand.Category() == consts.FullHouse {
		return false
	} else if madeHand.Category() != consts.Flush {
		return true
	} else {
		anotherHand := madeHand.(Flush)
		fValue := f.Hand[4].Rank*10000 + f.Hand[3].Rank*1000 + f.Hand[2].Rank*100 + f.Hand[1].Rank*10 + f.Hand[0].Rank
		aValue := anotherHand.Hand[4].Rank*10000 + anotherHand.Hand[3].Rank*1000 + anotherHand.Hand[2].Rank*100 + anotherHand.Hand[1].Rank*10 + anotherHand.Hand[0].Rank
		return fValue-aValue > 0
	}
}
