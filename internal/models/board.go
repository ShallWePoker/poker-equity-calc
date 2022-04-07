package models

import (
	"errors"
	"fmt"
)

type Board []Card

func (board Board) ToString() string {
	str := ""
	for _, card := range board {
		str = str + card.ToString()
	}
	return str
}

func InitBoard(cards []Card) (Board, error) {
	if len(cards) < 3 || len(cards) > 5 {
		return Board{}, errors.New("board must have 3 to 5 cards")
	}
	err := examineDistinctCards(cards...)
	if err != nil {
		return Board{}, err
	}
	return cards, nil
}

func examineDistinctCards(cards ...Card) error {
	distinctCards := make(map[Card]struct{})
	for _, card := range cards {
		_, exists := distinctCards[card]
		if !exists {
			distinctCards[card] = struct{}{}
		} else {
			return errors.New(fmt.Sprintf("duplicate card: %s", card.ToString()))
		}
	}
	return nil
}
