package models

import "fmt"

type HoleCardsEquity struct {
	HoleCards HoleCards
	Equity    float64
	WinRate   float64
	TieRate   float64
}

func (hce HoleCardsEquity) ToString() string {
	return fmt.Sprintf("holecard %s equity: Equity %.3f%%, WinRate %.3f%%, TieRate %.3f%%", hce.HoleCards.ToString(), hce.Equity*100, hce.WinRate*100, hce.TieRate*100)
}
