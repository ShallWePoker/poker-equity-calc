package models

import "github.com/ShallWePoker/poker-equity-calc/internal/consts"

type FullHouse struct {
	Hand          Hand
	ThreeCardRank int
	TwoCardRank   int
}

func (f FullHouse) ToString() string {
	return f.Hand.ToString()
}

func (f FullHouse) BoardCards() Hand {
	return f.Hand
}

func (f FullHouse) Category() string {
	return consts.FullHouse
}

func (f FullHouse) IsGreaterThan(madeHand MadeHand) bool {
	if madeHand.Category() == consts.StraightFlush || madeHand.Category() == consts.FourOfAKind {
		return false
	} else if madeHand.Category() != consts.FullHouse {
		return true
	} else {
		anotherHand := madeHand.(FullHouse)
		if f.ThreeCardRank != anotherHand.ThreeCardRank {
			return f.ThreeCardRank - anotherHand.ThreeCardRank > 0
		} else {
			return f.TwoCardRank - anotherHand.TwoCardRank > 0
		}
	}
}