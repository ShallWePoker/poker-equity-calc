package models

import (
	"errors"
	"strings"
)

type Flop []Card

func (Flop Flop) ToString() string {
	face := ""
	for _, card := range Flop {
		face = face + card.ToString()+" "
	}
	face = strings.TrimSuffix(face, " ")
	return face
}

func InitFlop(cards []Card) (Flop, error) {
	if len(cards) != 3 {
		return Flop{}, errors.New("flop must have 3 cards")
	}
	flop := Flop(cards)
	if flop[0] == flop[1] || flop[1] == flop[2] || flop[0] == flop[2] {
		return Flop{}, errors.New("flop must have 3 distinct cards")
	}
	return flop, nil
}
