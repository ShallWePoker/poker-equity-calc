package models

type MadeHand interface {
	BoardCards() Hand
	ToString() string
	IsGreaterThan(hand MadeHand) bool
	Category() string
}
