package models

import (
	"errors"
	"strings"
)

type HoleCards []Card

func (holeCards HoleCards) ToString() string {
	face := ""
	for _, card := range holeCards {
		face = face + card.ToString()+" "
	}
	face = strings.TrimSuffix(face, " ")
	return face
}

func InitHoleCards(cards []Card) (HoleCards, error) {
	if len(cards) != 2 {
		return HoleCards{}, errors.New("hole cards must have 2 cards")
	}
	hc := HoleCards(cards)
	if hc[0] == hc[1] {
		return HoleCards{}, errors.New("hole cards must have 2 distinct cards")
	}
	return hc, nil
}