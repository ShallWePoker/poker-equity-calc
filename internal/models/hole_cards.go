package models

import "strings"

type HoleCards []Card

func (holeCards HoleCards) ToString() string {
	face := ""
	for _, card := range holeCards {
		face = face + card.ToString()+" "
	}
	face = strings.TrimSuffix(face, " ")
	return face
}
