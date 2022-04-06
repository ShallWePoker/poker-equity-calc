package models

import "fmt"

type HoleCardsEquity struct {
	HoleCards HoleCards
	WinRate float64
	TieRate float64
}

func (hce HoleCardsEquity) ToString() string {
	return fmt.Sprintf("hole cards %s equity: WinRate %v, TieRate %v", hce.HoleCards.ToString(), hce.WinRate, hce.TieRate)
}
