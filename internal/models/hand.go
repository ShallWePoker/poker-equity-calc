package models

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"strings"
)

type Hand []Card

func InitHand(cards []Card) (Hand, error) {
	hand := make([]Card, 5)
	err := checkHandLength(cards)
	if err != nil {
		return hand, err
	}
	copy(hand, cards)
	sortHandByRank(hand)
	return hand, nil
}

func checkHandLength(hand []Card) error {
	if len(hand) != 5 {
		return errors.New(fmt.Sprintf("hand must have 5 cards"))
	}
	return nil
}

func (hand Hand) ToString() string {
	title := ""
	for _, card := range hand {
		title = title + card.Title + card.Suit + " "
	}
	return strings.TrimSuffix(title, " ")
}

func (hand Hand) ToRankString() string {
	title := ""
	for _, card := range hand {
		title = title + card.Title
	}
	return title
}

func sortHandByRank(hand Hand) {
	for i := 1; i < len(hand); i++ {
		tmp := hand[i]
		k := i
		for (k > 0) && tmp.Rank < hand[k-1].Rank {
			hand[k] = hand[k-1]
			k = k - 1
		}
		hand[k] = tmp
	}
}

func (hand Hand) IsFlush() bool {
	firstSuit := hand[0].Suit
	for _, card := range hand[1:len(hand)] {
		if card.Suit != firstSuit {
			return false
		}
	}
	return true
}

func (hand Hand) IsStraight() bool {
	return consts.Straights[hand.ToRankString()]
}

func (hand Hand) IsStraightFlush() bool {
	if hand.IsFlush() && hand.IsStraight() {
		return true
	}
	return false
}

func (hand Hand) IsFourOfAKind() bool {
	ctr := 0
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank == hand[i+1].Rank {
			for i < 4 && hand[i].Rank == hand[i+1].Rank {
				ctr = ctr + 1
				i = i + 1
			}
			i = i - 1
			if ctr == 3 {
				return true
			} else {
				ctr = 0
			}
		}
	}
	return false
}

func (hand Hand) IsFullHouse() bool {
	variation := 0
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank != hand[i+1].Rank {
			variation = variation + 1
		}
	}
	if variation == 1 {
		ctr := 0
		i := 0
		for hand[i].Rank == hand[i+1].Rank {
			ctr = ctr + 1
			i = i + 1
		}
		if ctr == 1 || ctr == 2 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (hand Hand) IsThreeOfAKind() bool {
	ctr := 1
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank == hand[i+1].Rank {
			for i < 4 && hand[i].Rank == hand[i+1].Rank {
				ctr = ctr + 1
				i = i + 1
			}
			if ctr == 3 {
				return true
			} else {
				ctr = 1
			}
		}
	}
	return false
}

func (hand Hand) IsTwoPairs() bool {
	pairs := 0
	for i := range []int{0, 1, 2} {
		if i < 3 && hand[i].Rank == hand[i+1].Rank && hand[i].Rank != hand[i+2].Rank {
			pairs = pairs + 1
		}
	}
	if hand[4].Rank == hand[3].Rank && hand[3].Rank != hand[2].Rank {
		pairs = pairs + 1
	}
	if pairs == 2 {
		return true
	} else {
		return false
	}
}

func (hand Hand) IsOnePair() bool {
	pairs := 0
	for i := range []int{0, 1, 2} {
		if i < 3 && hand[i].Rank == hand[i+1].Rank && hand[i].Rank != hand[i+2].Rank {
			pairs = pairs + 1
		}
	}
	if hand[4].Rank == hand[3].Rank && hand[3].Rank != hand[2].Rank {
		pairs = pairs + 1
	}
	if pairs == 1 {
		return true
	} else {
		return false
	}
}

func (hand Hand) Categorize() MadeHand {
	if hand.IsFlush() && hand.IsStraight() {
		madeHand := StraightFlush{
			Hand: hand,
		}
		return madeHand
	}
	if hand.IsFourOfAKind() {
		madeHand := FourOfAKind{
			Hand: hand,
		}
		if hand[0] == hand[1] {
			madeHand.FourCardRank = hand[0].Rank
			madeHand.OneCardRank = hand[1].Rank
		} else {
			madeHand.FourCardRank = hand[1].Rank
			madeHand.OneCardRank = hand[0].Rank
		}
		return madeHand
	}
	if hand.IsFullHouse() {
		madeHand := FullHouse{Hand: hand, ThreeCardRank: hand[2].Rank}
		if hand[2] == hand[0] {
			madeHand.TwoCardRank = hand[4].Rank
		} else {
			madeHand.TwoCardRank = hand[0].Rank
		}
		return madeHand
	}

	if hand.IsFlush() && !hand.IsStraight() {
		madeHand := Flush{Hand: hand}
		return madeHand
	}

	if !hand.IsFlush() && hand.IsStraight() {
		madeHand := Straight{Hand: hand}
		return madeHand
	}

	if !hand.IsFullHouse() && hand.IsThreeOfAKind() {
		madeHand := ThreeOfAKind{Hand: hand, ThreeCardRank: hand[2].Rank}
		if hand[0] == hand[1] {
			madeHand.HighCardRank = hand[4].Rank
			madeHand.LowCardRank = hand[3].Rank
		} else if hand[1] == hand[2] {
			madeHand.HighCardRank = hand[4].Rank
			madeHand.LowCardRank = hand[0].Rank
		} else {
			madeHand.HighCardRank = hand[1].Rank
			madeHand.LowCardRank = hand[0].Rank
		}
		return madeHand
	}

	if !hand.IsFullHouse() && hand.IsTwoPairs() {
		madeHand := TwoPairs{Hand: hand}
		if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[3].Rank {
			madeHand.LowTwoCardRank = hand[0].Rank
			madeHand.HighTwoCardRank = hand[2].Rank
			madeHand.SingleCardRank = hand[4].Rank
		} else if hand[1].Rank == hand[2].Rank {
			madeHand.LowTwoCardRank = hand[1].Rank
			madeHand.HighTwoCardRank = hand[4].Rank
			madeHand.SingleCardRank = hand[0].Rank
		} else {
			madeHand.LowTwoCardRank = hand[0].Rank
			madeHand.HighTwoCardRank = hand[4].Rank
			madeHand.SingleCardRank = hand[2].Rank
		}
		return madeHand
	}

	if !hand.IsFullHouse() && !hand.IsThreeOfAKind() && !hand.IsTwoPairs() && hand.IsOnePair() {
		madeHand := OnePair{Hand: hand}
		if hand[0].Rank == hand[1].Rank {
			madeHand.PairCardRank = hand[0].Rank
			madeHand.HighCardRank = hand[4].Rank
			madeHand.MidCardRank = hand[3].Rank
			madeHand.LowCardRank = hand[2].Rank
		} else if hand[1].Rank == hand[2].Rank {
			madeHand.PairCardRank = hand[1].Rank
			madeHand.HighCardRank = hand[4].Rank
			madeHand.MidCardRank = hand[3].Rank
			madeHand.LowCardRank = hand[0].Rank
		} else if hand[2].Rank == hand[3].Rank {
			madeHand.PairCardRank = hand[2].Rank
			madeHand.HighCardRank = hand[4].Rank
			madeHand.MidCardRank = hand[1].Rank
			madeHand.LowCardRank = hand[0].Rank
		} else {
			madeHand.PairCardRank = hand[4].Rank
			madeHand.HighCardRank = hand[2].Rank
			madeHand.MidCardRank = hand[1].Rank
			madeHand.LowCardRank = hand[0].Rank
		}
		return madeHand
	}

	madeHand := HighCard{Hand: hand}

	return madeHand
}
