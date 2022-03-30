package models

import (
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"strconv"
)

type Flush struct {
	Hand Hand
}

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
		fValue, _ := strconv.Atoi(f.Hand.ToRankString())
		aValue, _ := strconv.Atoi(anotherHand.Hand.ToRankString())
		return fValue - aValue > 0
	}
}