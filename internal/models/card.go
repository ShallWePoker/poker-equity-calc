package models

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"strings"
)

var ValidSuits = map[string]string{
	"H": "❤",
	"D": "♦",
	"S": "♠",
	"C": "♣",
}

var ValidTitles = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var cardsStrEnums = []string{
	"2H", "2S", "2D", "2C",
	"3H", "3S", "3D", "3C",
	"4H", "4S", "4D", "4C",
	"5H", "5S", "5D", "5C",
	"6H", "6S", "6D", "6C",
	"7H", "7S", "7D", "7C",
	"8H", "8S", "8D", "8C",
	"9H", "9S", "9D", "9C",
	"TH", "TS", "TD", "TC",
	"JH", "JS", "JD", "JC",
	"QH", "QS", "QD", "QC",
	"KH", "KS", "KD", "KC",
	"AH", "AS", "AD", "AC",
}

var cardsEnums = make([]Card, 0)

func InitCardsEnums() {
	for _, str := range cardsStrEnums {
		card, _ := InitCardFromString(str)
		cardsEnums = append(cardsEnums, card)
	}
}

func GetCardsEnums() []Card {
	return cardsEnums
}

type Card struct {
	Title string `json:"title"`
	Suit  string `json:"suit"`
	Rank  int    `json:"rank"`
}

func InitCardFromString(str string) (Card, error) {
	if len(str) != 2 {
		return Card{}, errors.New("card must have title and suit; invalid input: " + str)
	}
	card := Card{Title: string(str[0]), Suit: string(str[1])}
	if err := card.format(); err != nil {
		return Card{}, err
	}
	return card, nil
}

func (card Card) ToString() string {
	return card.Title + card.Suit
}

func (card Card) ToStringValidSuits() string {
	return card.Title + consts.SuitsValid[card.Suit]
}

func (card *Card) format() error {
	card.Title = strings.ToUpper(card.Title)
	card.Suit = strings.ToUpper(card.Suit)
	var ok bool
	card.Rank, ok = ValidTitles[card.Title]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid input card: %s ", card.Title))
	}
	card.Suit, ok = ValidSuits[card.Suit]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid input card: %s ", card.Suit))
	}
	return nil
}
