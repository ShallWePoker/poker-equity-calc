package models

import (
	"errors"
	"strings"
)

type HoleCard []Card

func (holeCard HoleCard) ToString() string {
	face := ""
	for _, card := range holeCard {
		face = face + card.ToString()+" "
	}
	face = strings.TrimSuffix(face, " ")
	return face
}

func InitHoleCard(cards []Card) (HoleCard, error) {
	if len(cards) != 2 {
		return HoleCard{}, errors.New("hole cards must have 2 cards")
	}
	hc := HoleCard(cards)
	if hc[0] == hc[1] {
		return HoleCard{}, errors.New("hole cards must have 2 distinct cards")
	}
	return hc, nil
}

func (holeCard HoleCard) Key() string {
	face := ""
	for _, card := range holeCard {
		face = face + card.ToStringValidSuits()+" "
	}
	face = strings.TrimSuffix(face, " ")
	return face
}

/*
func holecardBlockedBy(holecard models.HoleCard, excludedCards ...models.Card) bool {
	for _, blockCard := range excludedCards {
		if holecard.BlockedBy(blockCard) {
			return true
		}
	}
	return false
}
 */

func (holeCard HoleCard) BlockedBy(cards ...Card) bool {
	for _, card := range cards {
		if holeCard[0] == card || holeCard[1] == card {
			return true
		}
	}
	return false
}