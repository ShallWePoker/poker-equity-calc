package models

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"strings"
)

type Card struct {
	Title string   `json:"title"`
	Suit  string   `json:"suit"`
	Rank  int      `json:"rank"`
}

func InitCardFromString(str string) (Card, error) {
	if len(str) != 2 {
		return Card{}, errors.New("card must have title and suit; invalid input: "+str)
	}
	card := Card{Title: string(str[0]), Suit: string(str[1])}
	if err := card.format(); err != nil {
		return Card{}, err
	}
	return card, nil
}

func (card Card) ToString() string {
	return card.Title+card.Suit
}

func (card *Card) format() error {
	card.Title = strings.ToUpper(card.Title)
	card.Suit = strings.ToUpper(card.Suit)
	var ok bool
	card.Rank, ok = consts.ValidTitles[card.Title]
	card.Suit, ok = consts.ValidSuits[card.Suit]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid input card: %s or %s", card.Title, card.Suit))
	}
	return nil
}