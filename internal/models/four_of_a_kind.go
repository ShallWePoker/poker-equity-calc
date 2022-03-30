package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type FourOfAKind struct {
	Hand         Hand
	FourCardRank int
	OneCardRank  int
}

func (f FourOfAKind) ToString() string {
	return f.Hand.ToString()
}

func (f FourOfAKind) BoardCards() Hand {
	return f.Hand
}

func (f FourOfAKind) Category() string {
	return consts.FourOfAKind
}

func (f FourOfAKind) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush {
		return false
	} else if madeHand.Category() != consts.FourOfAKind {
		return true
	} else {
		anotherHand := madeHand.(FourOfAKind)
		if f.FourCardRank != anotherHand.FourCardRank {
			return f.FourCardRank - anotherHand.FourCardRank > 0
		} else {
			return f.OneCardRank - anotherHand.OneCardRank > 0
		}
	}
}