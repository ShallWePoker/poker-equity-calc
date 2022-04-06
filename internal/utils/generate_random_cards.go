package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"math/rand"
	"time"
)

func removeSpecifiedCards(cards []models.Card) ([]models.Card, error) {
	backup := make([]models.Card, len(models.GetCardsEnums()))
	copy(backup, models.GetCardsEnums())
	for _, card := range cards {
		index := findIndexOfElement(backup, card)
		if index == -1 {
			return []models.Card{}, errors.New(fmt.Sprintf("card %s not found.", card))
		}
		backup = removeKthElement(backup, index)
	}
	return backup, nil
}

func findIndexOfElement(s []models.Card, target models.Card) int {
	for i, elem := range s {
		if elem == target {
			return i
		}
	}
	return -1
}

func removeKthElement(s []models.Card, k int) []models.Card {
	s[k] = s[len(s)-1]
	return s[:len(s)-1]
}


func GenerateRandomNCards(n int, removedCards []models.Card) ([]models.Card, error) {
	if n < 1 || n > 7 {
		return []models.Card{}, errors.New("n must be in [1, 7]")
	}
	cardsLeft, err := removeSpecifiedCards(removedCards)
	if err != nil {
		return []models.Card{}, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cardsLeft), func (i, j int) {
		cardsLeft[i], cardsLeft[j] = cardsLeft[j], cardsLeft[i]
	})
	cards := make([]models.Card, 0)
	for i := 0; i < n ; i++ {
		cards = append(cards, cardsLeft[i])
	}
	return cards, nil
}
