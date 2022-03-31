package utils

import (
	"errors"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
	"math/rand"
	"time"
)

func removeSpecifiedCards(cards []string) []string {
	backup := make([]string, len(consts.CardsEnums))
	copy(backup, consts.CardsEnums)
	for _, card := range cards {
		index := findIndexOfElement(backup, card)
		backup = removeKthElement(backup, index)
	}
	return backup
}

func findIndexOfElement(s []string, target string) int {
	for i, elem := range s {
		if elem == target {
			return i
		}
	}
	return 0
}

func removeKthElement(s []string, k int) []string {
	s[k] = s[len(s)-1]
	return s[:len(s)-1]
}


func GenerateRandomNCards(n int, removedCards []string) ([]models.Card, error) {
	if n < 1 || n > 7 {
		return []models.Card{}, errors.New("n must be in [1, 7]")
	}
	cardsLeft := removeSpecifiedCards(removedCards)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cardsLeft), func (i, j int) {
		cardsLeft[i], cardsLeft[j] = cardsLeft[j], cardsLeft[i]
	})
	cards := make([]models.Card, 0)
	for i := 0; i < n ; i++ {
		card, err := models.InitCardFromString(cardsLeft[i])
		if err != nil {
			return []models.Card{}, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
